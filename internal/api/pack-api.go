package api

import (
	"better-shipping-app/internal/service"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type packApi struct {
	service service.PackService
}

func (p packApi) registerRoutes(s Server) {
	s.getMux().HandleFunc("GET /packs", p.listPacks)
}

func (p packApi) listPacks(writer http.ResponseWriter, _ *http.Request) {
	list, err := p.service.ListPacks()

	if err != nil {
		log.Error(err)
		handleHttpError(writer, 500, "Internal Server Error")
		return
	}

	respondJsonBody(writer, list)
}

func RegisterPackApi(packService service.PackService, server Server) {
	p := &packApi{service: packService}

	p.registerRoutes(server)
}
