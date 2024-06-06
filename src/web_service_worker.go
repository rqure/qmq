package qmq

import (
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WebServiceWorkerSignals struct {
	ClientConnected    Signal
	ClientDisconnected Signal
	Received           Signal
}

type WebServiceWorker struct {
	clients map[string]IWebClient
	signals WebServiceWorkerSignals
}

func NewWebServiceWorker() *WebServiceWorker {
	return &WebServiceWorker{}
}

func (w *WebServiceWorker) Init() {
	// Serve static files from the "static" directory
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./web/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./web/img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./web/js"))))

	// Handle WebSocket and other routes
	http.Handle("/", w)

	// Register_web_handler_notification_listener()
	// Register_web_handler_notification_manager()
	// Register_web_handler_server_interactor()
	// Register_web_handler_app()

	go func() {
		err := http.ListenAndServe("0.0.0.0:20000", nil)
		if err != nil {
			Panic("[WebServiceWorker::Init] HTTP server error: %v", err)
		}
	}()
}

func (w *WebServiceWorker) onIndexRequest(wr http.ResponseWriter, _ *http.Request) {
	index, err := os.ReadFile("web/index.html")

	if err != nil {
		Error("[WebServiceWorker::onIndexRequest] Error reading index.html: %v", err)
		return
	}

	wr.Header().Set("Content-Type", "text/html")
	wr.Write(index)
}

func (w *WebServiceWorker) onWSRequest(wr http.ResponseWriter, req *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(wr, req, nil)
	if err != nil {
		Error("[WebServiceWorker::onWSRequest] Error upgrading to WebSocket: %v", err)
		return
	}

	w.addClient(conn)
}

func (w *WebServiceWorker) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		w.onIndexRequest(wr, req)
	} else if req.URL.Path == "/ws" {
		w.onWSRequest(wr, req)
	} else {
		http.NotFound(wr, req)
	}
}

func (w *WebServiceWorker) Deinit() {
	for _, client := range w.clients {
		client.Close()
	}
}

func (w *WebServiceWorker) DoWork() {
	for _, client := range w.clients {
		if m := client.Read(); m != nil {
			w.signals.Received.Emit(client, m)
		}
	}
}

func (w *WebServiceWorker) Send(clientId string, p *anypb.Any) {
	if client, ok := w.clients[clientId]; ok {
		client.Write(&WebMessage{
			Header: &WebHeader{
				Id:        uuid.New().String(),
				Timestamp: timestamppb.Now(),
			},
			Payload: p,
		})
	}
}

func (w *WebServiceWorker) Broadcast(p *anypb.Any) {
	for clientId := range w.clients {
		w.Send(clientId, p)
	}
}

func (w *WebServiceWorker) addClient(conn *websocket.Conn) IWebClient {
	client := NewWebClient(conn, func(id string) {
		w.signals.ClientDisconnected.Emit(id)
		delete(w.clients, id)
	})
	w.clients[client.Id()] = client

	w.signals.ClientConnected.Emit(client)

	return client
}