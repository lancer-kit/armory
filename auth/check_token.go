package auth

import (
	"net/http"

	"gitlab.inn4science.com/vcg/go-common/api/render"
	"gitlab.inn4science.com/vcg/go-common/log"
	"github.com/pkg/errors"
	"encoding/json"
	"context"
)

// Header name of the `Authorization` header.
const Header = "Authorization"
const JWTHeader = "jwt"
var userApiLink string
const KeyUID = iota
func Init(usrApiLink string ) {
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

func ExtractUserID() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rawJwt := r.Header.Get(JWTHeader)
			if rawJwt == "" {
				render.ResultBadRequest.
					SetError("JWT Header must not bee empty").
					Render(w)
				return
			}
			jwt := struct {
				Jti int64 `json:"jti,string"`
			}{}
			err := json.Unmarshal([]byte(rawJwt), &jwt)
			if err != nil {
				render.ResultBadRequest.
					SetError("JWT Header is invalid json").
					Render(w)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), KeyUID, jwt.Jti))
			next.ServeHTTP(w, r)
			return

		})
	}
}
