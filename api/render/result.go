package render

import (
	"encoding/json"
	"net/http"
)

// PrettyMarshal is a flag that enable marshalling with indent
var PrettyMarshal bool

// WriteJSON writes some response as WriteJSON to the `http.ResponseWriter`.
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	var marshaled []byte
	var err error

	if PrettyMarshal {
		marshaled, err = json.MarshalIndent(data, "", "  ")
	} else {
		marshaled, err = json.Marshal(data)
	}

	if err != nil {
		http.Error(w, "error while render response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(marshaled)
}

// ServerError renders default http.StatusInternalServerError.
func ServerError(w http.ResponseWriter) {
	WriteJSON(w, http.StatusInternalServerError, ResultServerError)
}
