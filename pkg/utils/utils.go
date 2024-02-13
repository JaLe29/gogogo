package utils

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("dssdadsasdasdasda")

func ReadUserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Original-Forwarded-For")
	if IPAddress != "" {
		return IPAddress
	}

	IPAddress = r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func GenerateJwtToken() string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "my-auth-server",
		"sub": "john",
		"foo": 2,
	})
	s, e := t.SignedString(secret)

	if e != nil {
		panic(e)
	}

	return s
}
