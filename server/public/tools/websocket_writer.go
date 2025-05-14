package tools

import (
	"github.com/gorilla/websocket"
)

// WSWriter 封装了一个后台写协程，串行化所有写操作
type WSWriter struct {
	ws   *websocket.Conn
	ch   chan []byte
	done chan struct{}
}

// NewChanWriter 构造一个仅用 channel 的 WSWriter，适用于内部推消息
func NewChanWriter(ch chan []byte) *WSWriter {
	return &WSWriter{
		ws:   nil,
		ch:   ch,
		done: make(chan struct{}),
	}
}

// NewWSWriter 构造并启动后台写协程
func NewWSWriter(ws *websocket.Conn) *WSWriter {
	w := &WSWriter{
		ws:   ws,
		ch:   make(chan []byte, 1000),
		done: make(chan struct{}),
	}
	go w.writeLoop()
	return w
}
func (w *WSWriter) writeLoop() {
	for {
		select {
		case msg, ok := <-w.ch:
			if !ok {
				return
			}
			if w.ws != nil {
				// 只有 WebSocket 模式才真正发送
				w.ws.WriteMessage(websocket.TextMessage, msg)
			}
		case <-w.done:
			return
		}
	}
}

// Send 将消息排入后台写队列
func (w *WSWriter) Send(msg []byte) {
	select {
	case w.ch <- msg:
	default:
		// 通道已满时丢弃，防止阻塞
	}
}

// Close 停止后台写协程并关闭底层连接
func (w *WSWriter) Close() {
	close(w.done)
	w.ws.Close()
}
