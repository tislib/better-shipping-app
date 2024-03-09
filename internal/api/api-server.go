package api

import (
	"better-shipping-app/internal/config"
	"github.com/rs/cors"
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
	c := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		MaxAge:           86400,
	})

	return http.ListenAndServe(s.config.ADDR, c.Handler(s.mux))
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
