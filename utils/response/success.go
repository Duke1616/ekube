package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Success(w http.ResponseWriter, data interface{}) {
	c := 0
	resp := Data{
		Code:    &c,
		Message: "",
		Data:    data,
	}

	// set response heanders
	w.Header().Set("Content-Type", "application/json")

	respBytes, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMSG := fmt.Sprintf(`{"status":"error", "message": "encoding to json error, %s"}`, err)
		w.Write([]byte(errMSG))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
}
