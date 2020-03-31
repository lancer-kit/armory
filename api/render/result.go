package render

import (
	"net/http"
)

// R is a default struct for json responses.

//Example:
//
//``` go
//	func MyHandler(w http.ResponseWriter, r *http.Request) {
//		// some code ...
//		// ...
//		res := render.R{
//			Code: http.StatusOk,
//			Message: "User created",
//		}
//		res.Render(w)
//		return
//	}
//```
// Usage of predefined response:
//``` go
//	func MyHandler(w http.ResponseWriter, r *http.Request) {
//		// some code ...
//		// ...
//		render.ResultBadRequest.SetError("Invalid email").Render(w)
//		return
//	}
//```
type R struct {
	Code    int         `json:"errcode"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"errmsg,omitempty"`
}

// SetError adds error details to response.
func (r *R) SetError(val interface{}) *R {
	clone := *r

	switch casted := val.(type) {
	case nil:
		break
	case error:
		clone.Error = casted.Error()
	case string:
		clone.Error = casted
	case R:
		clone.Error = casted.Error
	case *R:
		clone.Error = casted.Error
	default:
		clone.Error = val
	}

	return &clone
}

// SetData sets response data.
func (r *R) SetData(val interface{}) *R {
	nr := *r
	nr.Data = val
	return &nr
}

// Render writes current response as WriteJSON to the `http.ResponseWriter`.
func (r *R) Render(w http.ResponseWriter) {
	WriteJSON(w, r.Code, r)
}
