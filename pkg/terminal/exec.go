package terminal

import (
	"github.com/gorilla/websocket"
	"io"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

type ttyHandler interface {
	Stdin() io.Reader
	Stdout() io.Writer
	Stderr() io.Writer
	Tty() bool
	remotecommand.TerminalSizeQueue
	Close()
}

type sizeQueue struct {
	resizeChan   chan remotecommand.TerminalSize
	stopResizing chan struct{}
}

var _ remotecommand.TerminalSizeQueue = &sizeQueue{}

func (s *sizeQueue) Next() *remotecommand.TerminalSize {
	select {
	case size := <-s.resizeChan:
		return &size
	case <-s.stopResizing:
		return nil
	}
}

type Terminal struct {
	client kubernetes.Interface
	config *rest.Config
	option *Option
}

func NewTerminal(client kubernetes.Interface, config *rest.Config, option *Option) Interface {
	return &Terminal{client: client, config: config, option: option}
}

func (t *Terminal) Exec(cmd []string, namespace, pod, container string, handler ttyHandler) error {
	//pods, _ := t.client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	//fmt.Print(pods)
	req := t.client.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespace).
		Name(pod).
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		Container: container,
		Command:   cmd,
		Stdin:     handler.Stdin() != nil,
		Stdout:    handler.Stdout() != nil,
		Stderr:    handler.Stderr() != nil,
		TTY:       handler.Tty(),
	}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(t.config, "POST", req.URL())
	if err != nil {
		return err
	}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler.Stdin(),
		Stdout:            handler.Stdout(),
		Stderr:            handler.Stderr(),
		TerminalSizeQueue: handler,
		Tty:               handler.Tty(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (t *Terminal) HandleSession(shell, namespace, pod, container string, conn *websocket.Conn) {
	var err error
	validShells := []string{"bash", "sh"}

	handler := &kubeShell{conn: conn, sizeChan: make(chan remotecommand.TerminalSize)}

	if isValidShell(validShells, shell) {
		cmd := []string{shell}
		err = t.Exec(cmd, namespace, pod, container, handler)
	} else {
		for _, testShell := range validShells {
			cmd := []string{testShell}
			if err = t.Exec(cmd, namespace, pod, container, handler); err == nil {
				break
			}
		}
	}

	if err != nil {
		handler.Close()
		return
	}

	handler.Close()
}

func isValidShell(validShells []string, shell string) bool {
	for _, validShell := range validShells {
		if validShell == shell {
			return true
		}
	}
	return false
}
