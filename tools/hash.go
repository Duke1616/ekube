package tools

import (
	"fmt"
	"hash/fnv"
	"strings"
)

func GenHashID(service, grpcPath string) string {
	hashedStr := fmt.Sprintf("%s-%s", service, grpcPath)
	h := fnv.New32a()
	h.Write([]byte(hashedStr))
	return fmt.Sprintf("%x", h.Sum32())
}

func FnvHash(contents ...string) string {
	h := fnv.New64a()
	h.Write([]byte(strings.Join(contents, "")))
	return fmt.Sprintf("%x", h.Sum64())
}
