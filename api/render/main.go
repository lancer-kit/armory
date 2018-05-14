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

// Success renders `result` as JSON with `http.StatusOK`.
func Success(w http.ResponseWriter, result interface{}) {
	WriteJSON(w, http.StatusOK, result)
}

// BadRequest renders `ResultBadRequest` with `reason` as an error.
func BadRequest(w http.ResponseWriter, reason interface{}) {
	ResultBadRequest.SetError(reason).Render(w)
}

// Unauthorized renders `ResultUnauthorized` with `reason` as an error.
func Unauthorized(w http.ResponseWriter, reason interface{}) {
	ResultUnauthorized.SetError(reason).Render(w)
}
