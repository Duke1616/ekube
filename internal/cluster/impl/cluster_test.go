package impl_test

import (
	"ekube/internal/cluster"
	"encoding/json"
	"io"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"testing"
)

func TestCreateCluster(t *testing.T) {
	req := cluster.NewCreateClusterRequest()
	req.Provider = "私有云"
	req.Region = "北京"
	req.Name = "k8s-cluster"

	req.KubeConfig = MustReadContentFile(filepath.Join(homedir.HomeDir(), ".kube", "config"))

	ins, err := impl.CreateCluster(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(MustToJson(ins))
}

func MustToJson(v any) string {
	return Prettify(v)
}

func Prettify(i interface{}) string {
	resp, _ := json.MarshalIndent(i, "", "   ")
	return string(resp)
}

func MustReadContentFile(filepath string) string {
	content, err := ReadContentFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func ReadContentFile(filepath string) ([]byte, error) {
	fd, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	payload, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
