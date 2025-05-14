package tools

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

// NewSSHClient 与之前相同...
func NewSSHClient(conf *SSHClientConfig) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		Timeout:         conf.Timeout,
		User:            conf.UserName,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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
	return ssh.Dial("tcp", fmt.Sprintf("%s:%d", conf.PublicIP, conf.Port), config)
}

// Turn 封装了 SSH 会话和 WS 写队列
type Turn struct {
	TaskName  string // 新增：哪个子任务组
	NodeName  string // 新增：哪个节点
	StdinPipe io.WriteCloser
	Session   *ssh.Session
	WsConn    *websocket.Conn // 仅用于读取客户端输入
	writer    *WSWriter       // 后台唯一写者
	LogWriter io.Writer       // 可选，把所有输出也写到这里

}

// NewTurn 接收 WSWriter，并启动交互式 Shell
func NewTurn(wsConn *websocket.Conn, sshClient *ssh.Client, writer *WSWriter, nodeName string) (*Turn, error) {
	sess, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}
	stdinPipe, err := sess.StdinPipe()
	if err != nil {
		return nil, err
	}
	turn := &Turn{
		NodeName:  nodeName, // 新增：哪个节点
		StdinPipe: stdinPipe,
		Session:   sess,
		WsConn:    wsConn,
		writer:    writer,
	}
	// 将 stdout 和 stderr 都定向到 Turn.Write
	sess.Stdout = turn
	sess.Stderr = turn
	// 请求 PTY 并启动 Shell
	modes := ssh.TerminalModes{ssh.ECHO: 0, ssh.TTY_OP_ISPEED: 14400, ssh.TTY_OP_OSPEED: 14400}
	if err := sess.RequestPty("dumb", 150, 30, modes); err != nil {
		return nil, err
	}
	if err := sess.Shell(); err != nil {
		return nil, err
	}
	return turn, nil
}

// Write 实现 io.Writer，将输出投递给后台唯一写者
func (t *Turn) Write(p []byte) (n int, err error) {
	msg := struct {
		Node string `json:"node"`
		Data string `json:"data"`
	}{
		Node: t.NodeName,
		Data: string(p),
	}
	b, _ := json.Marshal(msg)
	t.writer.Send(b)
	// 2) 如果绑定了 LogWriter，就一并写入
	if t.LogWriter != nil {
		if _, err := t.LogWriter.Write(p); err != nil {
			// 写缓冲失败也不影响前端输出，所以只记个 log 就好
			fmt.Printf("[Turn %s@%s] 写 LogWriter 失败: %v\n", t.TaskName, t.NodeName, err)
		}
	}

	return len(p), nil
}

// Read 从 WebSocket 读取客户端输入
func (t *Turn) Read(p []byte) (n int, err error) {
	for {
		msgType, reader, err := t.WsConn.NextReader()
		if err != nil {
			return 0, err
		}
		if msgType != websocket.TextMessage {
			continue
		}
		return reader.Read(p)
	}
}

// LoopRead 持续将前端输入写入远端 Shell stdin
func (t *Turn) LoopRead(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return errors.New("LoopRead exit")
		default:
			_, wsData, err := t.WsConn.ReadMessage()
			if err != nil {
				return fmt.Errorf("reading webSocket message err:%s", err)
			}
			// 解码、写入 stdin（可根据实际需要调整解码逻辑）
			body, _ := base64.StdEncoding.DecodeString(string(wsData[1:]))
			if _, err := t.StdinPipe.Write(body); err != nil {
				return fmt.Errorf("StdinPipe write err:%s", err)
			}
		}
	}
}

// SessionWait 等待远端 Shell 结束
func (t *Turn) SessionWait() error {
	return t.Session.Wait()
}

// Close 只关闭 SSH 会话，WSWriter 在外部统一 Close
func (t *Turn) Close() error {
	return t.Session.Close()
}
