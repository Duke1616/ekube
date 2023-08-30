package terminal

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"k8s.io/client-go/tools/remotecommand"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// ctrl+d to close terminal.
	endOfTransmission = "\u0004"
)

type kubeShell struct {
	conn     *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	stopChan chan struct{}
	tty      bool
}

var _ ttyHandler = &kubeShell{}

func (k *kubeShell) Stdin() io.Reader {
	return k
}

func (k *kubeShell) Stdout() io.Writer {
	return k
}

func (k *kubeShell) Stderr() io.Writer {
	return k
}

func (k *kubeShell) Tty() bool {
	return true
}

func (k *kubeShell) Next() *remotecommand.TerminalSize {
	select {
	case size := <-k.sizeChan:
		return &size
	case <-k.stopChan:
		return nil
	}
}

func (k *kubeShell) Close() {
	//close(k.stopChan)
	k.conn.Close()
}

func (k *kubeShell) Read(p []byte) (n int, err error) {
	var msg TermMessage
	err = k.conn.ReadJSON(&msg)
	if err != nil {
		return copy(p, endOfTransmission), err
	}

	switch msg.Type {
	case "stdin":
		return copy(p, msg.Data), nil
	case "resize":
		k.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
		return 0, nil
	default:
		return copy(p, endOfTransmission), fmt.Errorf("unknown message type '%s'", msg.Type)
	}
}

func (k *kubeShell) Write(p []byte) (n int, err error) {
	msg, err := json.Marshal(TermMessage{
		Type: "stdout",
		Data: string(p),
	})
	if err != nil {
		return 0, err
	}
	//
	k.conn.SetWriteDeadline(time.Now().Add(writeWait))
	if err = k.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return 0, err
	}
	return len(p), nil
}

type TermMessage struct {
	Type string `json:"type"`
	Data string `json:"data,omitempty"`
	Rows uint16 `json:"rows,omitempty"`
	Cols uint16 `json:"cols,omitempty"`
}
