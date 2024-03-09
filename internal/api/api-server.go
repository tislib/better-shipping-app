package api

import (
	"better-shipping-app/internal/config"
	"net/http"
)

type Server interface {
	Start() error
	getMux() *http.ServeMux
}

type server struct {
	mux    *http.ServeMux
	config config.ServerConfig
}

func (s *server) Start() error {
	return http.ListenAndServe(s.config.ADDR, s.mux)
}

func (s *server) getMux() *http.ServeMux {
	return s.mux
}

func NewServer(config config.ServerConfig) Server {
	mux := http.NewServeMux()

	return &server{
		mux:    mux,
		config: config,
	}
}
