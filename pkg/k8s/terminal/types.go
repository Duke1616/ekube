package terminal

import "github.com/gorilla/websocket"

type Interface interface {
	HandleSession(shell, namespace, podName, containerName string, conn *websocket.Conn)
}
