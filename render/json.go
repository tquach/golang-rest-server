package render

import (
	"encoding/json"
	"log"
	"net/http"
)

// Constants defined for reuse
const (
	ContentTypeHeader = "Content-Type"
	JSONContentType   = "application/json; charset=utf-8"
)

// ErrorMsg is a custom message type for errors
type ErrorMsg struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// JSONError will marshal an error to the reesponse stream.
func JSONError(err error, status int, w http.ResponseWriter) {
	JSON(ErrorMsg{
		Error: err.Error(),
		Code:  status,
	}, status, w)
}

// JSON will render an object to JSON. Content-type will be set by default to "application/json; charset=utf8".
func JSON(msg interface{}, status int, w http.ResponseWriter) {
	w.Header().Add(ContentTypeHeader, JSONContentType)
	w.WriteHeader(status)
	out, err := json.Marshal(msg)
	if err != nil {
		log.Println("[render] json marshalling failed", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// no need to handle error here
	_, _ = w.Write(out)
}
