package qmq

import (
	"context"
	"os"
	"strconv"
)

type QMQApplication struct {
	conn      *QMQConnection
	producers map[string]*QMQProducer
	consumers map[string]*QMQConsumer
	logger    *QMQLogger
}

func NewQMQApplication(ctx context.Context, name string) *QMQApplication {
	addr := os.Getenv("QMQ_ADDR")
	if addr == "" {
		addr = "redis:6379"
	}
	
	password := os.Getenv("QMQ_PASSWORD")
	
	logLength, err := strconv.Atoi(os.Getenv("QMQ_LOG_LENGTH"))
	if err != nil {
		logLength = 100
	}
	
	conn := NewQMQConnection(addr, password)
	
	return &QMQApplication{
		conn: conn,
		logger: NewQMQLogger(ctx, name, conn, logLength)
	}
}

func (a *QMQApplication) Initialize(ctx context.Context) {
	a.conn.Connect(ctx)
	a.logger.Initialize(ctx)
	a.logger.Advise(ctx, "Application has started")
}

func (a *QMQApplication) Deinitialize(ctx context.Context) {
	a.logger.Advise(ctx, "Application has stopped")
	a.conn.Disconnect()
}

func (a *QMQApplication) Producer(key string) *QMQProducer {
	return a.producers[key]
}

func (a *QMQApplication) Consumer(key string) *QMQConsumer {
	return a.consumers[key]
}

func (a *QMQApplication) AddProducer(ctx context.Context, key string, length int64) {
	a.producers[key] = NewQMQProducer(ctx, key, conn, length)
}

func (a *QMQApplication) AddConsumer(ctx context.Context, key string) {
	s.consumers[key] = NewQMQConsumer(ctx, key, conn)
}
