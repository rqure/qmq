package qmq

import (
	"crypto/rand"
	"encoding/base64"
	"time"
)

type QMQLocker struct {
	id    string
	token string
	conn  *QMQConnection
}

func NewQMQLocker(id string, conn *QMQConnection) *QMQLocker {
	return &QMQLocker{
		id:    id,
		token: "",
		conn:  conn,
	}
}

func (l *QMQLocker) TryLockWithTimeout(timeoutMs int64) bool {
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return false
	}

	l.token = base64.StdEncoding.EncodeToString(randomBytes)

	writeRequest := NewWriteRequest(&QMQString{Value: l.token})

	result, err := l.conn.TempSet(l.id, writeRequest, timeoutMs)
	if err != nil {
		return false
	}

	return result
}

func (l *QMQLocker) TryLock() bool {
	return l.TryLockWithTimeout(30000)
}

func (l *QMQLocker) Lock() {
	for !l.TryLock() {
		time.Sleep(100 * time.Millisecond)
	}
}

func (l *QMQLocker) LockWithTimeout(timeoutMs int64) {
	for !l.TryLockWithTimeout(timeoutMs) {
		time.Sleep(100 * time.Millisecond)
	}
}

func (l *QMQLocker) Unlock() {
	readRequest, err := l.conn.Get(l.id)
	if err != nil {
		return
	}

	token := QMQString{}
	err = readRequest.Data.UnmarshalTo(&token)
	if err != nil {
		return
	}

	if l.token != token.Value {
		return
	}

	l.conn.Unset(l.id)
}
