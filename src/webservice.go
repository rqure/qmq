package qmq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"sync"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type WebSocketClient struct {
	readCh    chan map[string]interface{}
	writeCh   chan interface{}
	conn      *websocket.Conn
	connMutex sync.Mutex
	app       *QMQApplication
	wg        sync.WaitGroup
}

func NewWebSocketClient(conn *websocket.Conn, app *QMQApplication) *WebSocketClient {
	wsc := &WebSocketClient{
		readCh:  make(chan map[string]interface{}, 1),
		writeCh: make(chan interface{}, 1),
		conn:    conn,
		app:     app,
	}

	go wsc.DoPendingWrites()
	go wsc.DoPendingReads()

	return wsc
}

func (wsc *WebSocketClient) ReadJSON() chan map[string]interface{} {
	return wsc.readCh
}

func (wsc *WebSocketClient) WriteJSON(v interface{}) {
	wsc.writeCh <- v
}

func (wsc *WebSocketClient) Close() {
	wsc.connMutex.Lock()
	wsc.conn.Close()
	wsc.connMutex.Unlock()

	wsc.wg.Wait()
}

func (wsc *WebSocketClient) DoPendingReads() {
	wsc.wg.Add(1)
	defer wsc.wg.Done()

	for {
		wsc.connMutex.Lock()
		messageType, p, err := wsc.conn.ReadMessage()
		wsc.connMutex.Unlock()

		if err != nil {
			wsc.app.Logger().Error(fmt.Sprintf("Error reading WebSocket message: %v", err))
			break
		}

		if messageType == websocket.TextMessage {
			var data map[string]interface{}
			if err := json.Unmarshal(p, &data); err != nil {
				wsc.app.Logger().Error(fmt.Sprintf("Error decoding WebSocket message: %v", err))
				continue
			}

			wsc.readCh <- data
		}
	}

	close(wsc.writeCh)
	close(wsc.readCh)
}

func (wsc *WebSocketClient) DoPendingWrites() {
	wsc.wg.Add(1)
	defer wsc.wg.Done()

	for v := range wsc.writeCh {
		wsc.connMutex.Lock()
		if err := wsc.conn.WriteJSON(v); err != nil {
			wsc.app.Logger().Error(fmt.Sprintf("Error sending WebSocket message: %v", err))
		}
		wsc.connMutex.Unlock()
	}
}

type KeyValueResponse struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type DataUpdateResponse struct {
	Data KeyValueResponse `json:"data"`
}

type WebServiceContext interface {
	App() *QMQApplication
	Schema() interface{}
	NotifyClients(data interface{})
}

type WebServiceTickHandler interface {
	onTick(c WebServiceContext)
}

type WebServiceSetHandler interface {
	onSet(c WebServiceContext, key string, value interface{})
}

type WebService struct {
	clients      map[*websocket.Conn]*WebSocketClient
	clientsMutex sync.Mutex
	app          *QMQApplication
	schema       interface{}
	schemaMutex  sync.Mutex
	tickHandlers []WebServiceTickHandler
	setHandlers  []WebServiceSetHandler
}

func NewWebService() *WebService {
	return &WebService{
		clients: make(map[*websocket.Conn]*WebSocketClient),
		app:     NewQMQApplication("garage"),
	}
}

func (w *WebService) Initialize(schema interface{}) {
	w.app.Initialize()

	w.app.AddConsumer("garage:notifications:state").Initialize()
	w.app.AddConsumer("garage:notifications:shelly:connected").Initialize()

	// Serve static files from the "static" directory
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))

	// Handle WebSocket and other routes
	http.Handle("/", w)

	go func() {
		err := http.ListenAndServe("0.0.0.0:20000", nil)
		if err != nil {
			w.app.Logger().Panic(fmt.Sprintf("HTTP server error: %v", err))
		}
	}()

	w.schemaMutex.Lock()
	w.schema = schema
	w.app.Db().GetSchema(&w.schema)
	w.app.Db().SetSchema(&w.schema)
	w.schemaMutex.Unlock()
}

func (w *WebService) Deinitialize() {
	w.app.Deinitialize()
}

func (w *WebService) NotifyClients(data interface{}) {
	w.clientsMutex.Lock()
	defer w.clientsMutex.Unlock()
	for conn := range w.clients {
		w.clients[conn].WriteJSON(data)
	}
}

func (w *WebService) App() *QMQApplication {
	return w.app
}

func (w *WebService) Schema() interface{} {
	return w.schema
}

func (w *WebService) AddTickHandler(handler WebServiceTickHandler) {
	w.tickHandlers = append(w.tickHandlers, handler)
}

func (w *WebService) AddSetHandler(handler WebServiceSetHandler) {
	w.setHandlers = append(w.setHandlers, handler)
}

func (w *WebService) Tick() {
	w.schemaMutex.Lock()
	defer w.schemaMutex.Unlock()

	for _, handler := range w.tickHandlers {
		handler.onTick(w)
	}
}

func (w *WebService) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		w.onIndexRequest(wr, req)
	} else if req.URL.Path == "/ws" {
		w.onWSRequest(wr, req)
	} else {
		http.NotFound(wr, req)
	}
}

func (w *WebService) onIndexRequest(wr http.ResponseWriter, req *http.Request) {
	index, err := os.ReadFile("web/index.html")

	if err != nil {
		w.app.Logger().Error(fmt.Sprintf("Error while reading file for path '/': %v", err))
		return
	}

	wr.Header().Set("Content-Type", "text/html")
	wr.Write(index)
}

func (w *WebService) onWSRequest(wr http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(wr, req, nil)
	if err != nil {
		w.app.Logger().Error(fmt.Sprintf("Error upgrading to WebSocket: %v", err))
		return
	}
	defer conn.Close()

	client := w.addClient(conn)

	for data := range client.ReadJSON() {
		if cmd, ok := data["cmd"].(string); ok && cmd == "get" {
			if key, ok := data["key"].(string); ok {
				response := DataUpdateResponse{}

				schemaWrapper := reflect.ValueOf(&w.schema).Elem()
				schemaType := schemaWrapper.Type()

				for i := 0; i < schemaWrapper.NumField(); i++ {
					field := schemaWrapper.Field(i)
					tag := schemaType.Field(i).Tag.Get("qmq")

					if tag == key {
						response.Data.Key = key
						w.schemaMutex.Lock()
						response.Data.Value = reflect.ValueOf(field.Interface()).FieldByName("Value").Interface()
						w.schemaMutex.Unlock()

						if err := conn.WriteJSON(response); err != nil {
							w.app.Logger().Error(fmt.Sprintf("Error sending WebSocket message: %v", err))
						}

						break
					}
				}
			} else {
				w.app.Logger().Error("Invalid 'get' command: missing 'key'")
			}
		} else if cmd, ok := data["cmd"].(string); ok && cmd == "set" {
			if key, ok := data["key"].(string); ok {
				if value, ok := data["value"]; ok {
					response := DataUpdateResponse{}

					schemaWrapper := reflect.ValueOf(&w.schema).Elem()
					schemaType := schemaWrapper.Type()

					for i := 0; i < schemaWrapper.NumField(); i++ {
						field := schemaWrapper.Field(i)
						tag := schemaType.Field(i).Tag.Get("qmq")

						if tag == key {
							fieldValue := reflect.ValueOf(field.Interface())
							if fieldValue.Kind() != reflect.ValueOf(value).Kind() {
								w.app.Logger().Error("Invalid 'set' command: value type mismatch")
								break
							}

							w.schemaMutex.Lock()
							fieldValue.FieldByName("Value").Set(reflect.ValueOf(value))
							writeRequest := NewWriteRequest(field.Interface().(proto.Message))
							w.app.Db().Set(key, writeRequest)
							for _, handler := range w.setHandlers {
								handler.onSet(w, key, value)
							}
							w.schemaMutex.Unlock()

							response.Data.Key = key
							response.Data.Value = value

							w.NotifyClients(response)

							break
						}
					}
				} else {
					w.app.Logger().Error("Invalid 'set' command: missing 'value'")
				}
			} else {
				w.app.Logger().Error("Invalid 'set' command: missing 'key'")
			}
		} else {
			w.app.Logger().Error("Invalid WebSocket command")
		}
	}

	w.removeClient(conn)
}

func (w *WebService) addClient(conn *websocket.Conn) *WebSocketClient {
	w.clientsMutex.Lock()
	defer w.clientsMutex.Unlock()
	w.clients[conn] = NewWebSocketClient(conn, w.app)
	return w.clients[conn]
}

func (w *WebService) removeClient(conn *websocket.Conn) {
	w.clientsMutex.Lock()
	defer w.clientsMutex.Unlock()
	w.clients[conn].Close()
	delete(w.clients, conn)
}
