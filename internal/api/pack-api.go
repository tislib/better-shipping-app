package api

import (
	"better-shipping-app/internal/service"
	"net/http"
)

type packApi struct {
}

func (p packApi) registerRoutes(s Server) {
	s.getMux().HandleFunc("GET /packs", p.listPacks)
}

func (p packApi) listPacks(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("[]"))
}

func RegisterPackApi(packService service.PackService, server Server) {
	p := &packApi{}

	p.registerRoutes(server)
}
