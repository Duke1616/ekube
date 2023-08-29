package terminal

import (
	"io"
	"os"
	"testing"
)

type TestExec struct {
	tty bool
	sizeQueue
}

func (t *TestExec) Stdin() io.Reader {
	return os.Stdin
}

func (t *TestExec) Stdout() io.Writer {
	return os.Stdout
}

func (t *TestExec) Stderr() io.Writer {
	return os.Stderr
}

func (t *TestExec) Tty() bool {
	return t.tty
}

func (t *TestExec) Close() {
	close(t.stopResizing)
}

var _ ttyHandler = &TestExec{}

func TestNewTerminalOption(t *testing.T) {
	//var (
	//	err error
	//)
	//
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//// creates the clientset
	//
	//set, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	klog.Fatal(err)
	//}
	//
	//cmd := []string{
	//	"sh", "-c", "date",
	//}
	//
	//term := NewTerminal(set, config, nil)
	//
	//handler := &TestExec{
	//	tty: true,
	//	sizeQueue: sizeQueue{
	//		resizeChan:   make(chan remotecommand.TerminalSize),
	//		stopResizing: make(chan struct{}),
	//	},
	//}
	//
	//if err = term.Exec(cmd, "default", "nginx-54bcfc567b-h5rbh", "nginx", handler); err != nil {
	//	panic(err)
	//}

}
