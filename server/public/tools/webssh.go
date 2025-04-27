package tools

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

func NewSSHClient(conf *SSHClientConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         conf.Timeout,
		User:            conf.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), //忽略know_hosts检查
	}
	switch conf.AuthModel {
	case "password":
		config.Auth = []ssh.AuthMethod{ssh.Password(conf.Password)}
	case "privateKey":
		signer, err := ssh.ParsePrivateKey([]byte(conf.PrivateKey))
		if err != nil {
			return nil, err
		}
		config.Auth = []ssh.AuthMethod{ssh.PublicKeys(signer)}
	}
	c, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", conf.PublicIP, conf.Port), config)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Turn struct {
	StdinPipe io.WriteCloser
	Session   *ssh.Session
	WsConn    *websocket.Conn
	writeMu   sync.Mutex
}

func NewTurn(wsConn *websocket.Conn, sshClient *ssh.Client) (*Turn, error) {
	sess, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	stdinPipe, err := sess.StdinPipe()
	if err != nil {
		return nil, err
	}

	turn := &Turn{StdinPipe: stdinPipe, Session: sess, WsConn: wsConn}
	sess.Stdout = turn
	sess.Stderr = turn

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := sess.RequestPty("xterm", 150, 30, modes); err != nil {
		return nil, err
	}
	if err := sess.Shell(); err != nil {
		return nil, err
	}

	return turn, nil
}
func (t *Turn) Write(p []byte) (n int, err error) {
	t.writeMu.Lock()
	writer, err := t.WsConn.NextWriter(websocket.TextMessage)
	if err != nil {
		t.writeMu.Unlock()
		return 0, err
	}
	defer func() {
		writer.Close()
		t.writeMu.Unlock() // 写完毕后解锁
	}()
	fmt.Println("Write:" + string(p))
	return writer.Write(p)
}
func (t *Turn) Close() error {
	fmt.Println("Close()")
	if t.Session != nil {
		t.Session.Close()
	}

	return t.WsConn.Close()
}
func (t *Turn) Read(p []byte) (n int, err error) {
	for {
		msgType, reader, err := t.WsConn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.TextMessage {
			continue
		}
		fmt.Println("Write:" + string(p))
		return reader.Read(p)
	}
}
func (t *Turn) LoopRead(context context.Context) error {
	for {
		select {
		case <-context.Done():
			return errors.New("LoopRead exit")
		default:
			_, wsData, err := t.WsConn.ReadMessage()
			fmt.Println("本地输入：" + string(wsData))
			if err != nil {
				return fmt.Errorf("reading webSocket message err:%s", err)
			}
			body := decode(wsData[1:])
			fmt.Println("body:" + string(body))
			body = wsData
			fmt.Println("body:" + string(body))
			if _, err := t.StdinPipe.Write(body); err != nil {
				return fmt.Errorf("StdinPipe write err:%s", err)
			}

		}
	}
}

func (t *Turn) SessionWait() error {
	if err := t.Session.Wait(); err != nil {
		return err
	}
	return nil
}

func decode(p []byte) []byte {
	decodeString, _ := base64.StdEncoding.DecodeString(string(p))
	return decodeString
}

type WSWriter struct {
	ch   chan []byte
	ws   *websocket.Conn
	done chan struct{}
}

func NewWSWriter(ws *websocket.Conn) *WSWriter {
	w := &WSWriter{ws: ws, ch: make(chan []byte, 100), done: make(chan struct{})}
	go func() {
		for {
			select {
			case msg := <-w.ch:
				w.ws.WriteMessage(websocket.TextMessage, msg)
			case <-w.done:
				return
			}
		}
	}()
	return w
}

func (w *WSWriter) Send(msg []byte) {
	w.ch <- msg
}

func (w *WSWriter) Close() {
	close(w.done)
	w.ws.Close()
}
