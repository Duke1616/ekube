package api

import (
	terminal2 "ekube/pkg/k8s/terminal"
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"net/http"
)

var upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type terminalHandler struct {
	terminal terminal2.Interface
}

func newTerminalHandler(client kubernetes.Interface, config *rest.Config, option *terminal2.Option) *terminalHandler {
	return &terminalHandler{
		terminal: terminal2.NewTerminal(client, config, option),
	}
}

func (h *terminalHandler) handleTerminalSession(request *restful.Request, response *restful.Response) {
	namespace := request.PathParameter("namespace")
	podName := request.PathParameter("pod")
	//containerName := request.QueryParameter("container")
	//shell := request.QueryParameter("shell")

	conn, err := upgrade.Upgrade(response.ResponseWriter, request.Request, nil)
	if err != nil {
		klog.Warning(err)
		return
	}

	h.terminal.HandleSession("bash", namespace, podName, "nginx", conn)
}
