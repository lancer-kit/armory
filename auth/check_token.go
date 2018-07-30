package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"gitlab.inn4science.com/vcg/go-common/api/render"
	"gitlab.inn4science.com/vcg/go-common/log"
)

type ReturnAuthStruct struct {
	Jti     int64 `json:"jti,string"`
	IsAdmin bool  `json:"isAdmin,bool"`
}

// Header name of the `Authorization` header.
const (
	Header    = "Authorization"
	JWTHeader = "jwt"

	KeyUID     = "key_uid"
	KeyIsAdmin = "key_isAdmin"
)

var userApiLink string

func Init(usrApiLink string) {
	userApiLink = usrApiLink
}

// CheckToken checks `Authorization` token if it valid return nil.
func CheckToken(authtoken string) (int, []byte, error) {
	if userApiLink == "" {
		log.Default.Error("auth didn't init")
	}
	client := http.DefaultClient
	path := userApiLink + "/v1/auth"

	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return http.StatusInternalServerError,
			nil, errors.Wrap(err, "failed to create auth check request")
	}

	req.Header.Set(Header, authtoken)

	resp, err := client.Do(req)
	if err != nil {
		return http.StatusInternalServerError,
			nil, errors.Wrap(err, "failed to check token")
	}
	if resp.StatusCode == 200 {
		return 200, nil, nil
	}

	defer resp.Body.Close()
	respBody := make([]byte, 0)

	_, err = resp.Body.Read(respBody)
	if err != nil {
		log.Default.WithError(err).Error("unable to read response body")
		return http.StatusInternalServerError,
			nil, errors.Wrap(err, "failed read auth response body")
	}

	return resp.StatusCode, respBody, nil
}

// ValidateAuthHeader checks the request Authorization token.
// If token valid - continue request handling flow,
// else redirect `userapi` response to the requester.
func ValidateAuthHeader(required bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authtoken := r.Header.Get(Header)
			if authtoken == "" || !required {
				next.ServeHTTP(w, r)
				return
			}

			statusCode, rawResp, err := CheckToken(authtoken)
			if statusCode == http.StatusOK {
				next.ServeHTTP(w, r)
				return
			}

			if err != nil {
				log.Default.WithError(err).Error("unable to check auth token")
			}

			w.WriteHeader(statusCode)
			w.Write(rawResp)
		})
	}
}

// Method reads JWT Header and fill KeyUID and KeyIsAdmin in the context
// Use ExtractUserID() if jwt required
// Use ExtractUserID(false) if jwt not required
func ExtractUserID(required ...bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawJwt := r.Header.Get(JWTHeader)
			if rawJwt == "" {
				if len(required) == 0 || (len(required) > 0 && required[0]) {
					render.ResultBadRequest.
						SetError("JWT Header must not bee empty").
						Render(w)
					return
				}

				rCtx := context.WithValue(r.Context(), KeyUID, int64(0))
				rCtx = context.WithValue(rCtx, KeyIsAdmin, false)
				r = r.WithContext(rCtx)
				next.ServeHTTP(w, r)
				return
			}
			jwt := ReturnAuthStruct{}
			err := json.Unmarshal([]byte(rawJwt), &jwt)
			if err != nil {
				render.ResultBadRequest.
					SetError("JWT Header is invalid json").
					Render(w)
				return
			}

			rCtx := context.WithValue(r.Context(), KeyUID, jwt.Jti)
			rCtx = context.WithValue(rCtx, KeyIsAdmin, jwt.IsAdmin)
			r = r.WithContext(rCtx)
			next.ServeHTTP(w, r)
			return

		})
	}
}
