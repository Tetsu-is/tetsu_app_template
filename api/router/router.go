package router

import (
	"net/http"
	"time"
)

func NewServer(handler http.Handler) *http.Server {
	http.Handle("/", handler)
	s := &http.Server{Handler: handler, WriteTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second}
	return s
}
