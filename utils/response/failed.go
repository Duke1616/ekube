package response

import (
	"encoding/json"
	"fmt"
	"github.com/infraboard/mcube/exception"
	"net/http"
)

func Failed(w http.ResponseWriter, err error) {
	var (
		errCode  int
		httpCode int
		ns       string
		reason   string
		data     interface{}
		meta     interface{}
	)

	switch t := err.(type) {
	case exception.APIException:
		errCode = t.ErrorCode()
		reason = t.GetReason()
		data = t.GetData()
		meta = t.GetMeta()
		httpCode = t.GetHttpCode()
		ns = t.GetNamespace()
	default:
		errCode = exception.UnKnownException
	}

	if httpCode == 0 {
		httpCode = http.StatusInternalServerError
	}

	resp := Data{
		Code:      &errCode,
		Namespace: ns,
		Reason:    reason,
		Message:   err.Error(),
		Data:      data,
		Meta:      meta,
	}

	// set response heanders
	w.Header().Set("Content-Type", "application/json")

	// if marshal json error, use string to response
	respByt, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMSG := fmt.Sprintf(`{"status":"error", "message": "encoding to json error, %s"}`, err)
		w.Write([]byte(errMSG))
		return
	}

	w.WriteHeader(httpCode)
	w.Write(respByt)
}
