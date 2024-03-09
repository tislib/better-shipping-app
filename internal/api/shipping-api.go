package api

import (
	"better-shipping-app/internal/service"
	"net/http"
)

type shippingApi struct {
}

func (p shippingApi) registerRoutes(s Server) {
	s.getMux().HandleFunc("POST /shipping", p.calculate)
}

func (p shippingApi) calculate(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("[]"))
}

func RegisterShippingApi(shippingService service.ShippingService, server Server) {
	p := &shippingApi{}

	p.registerRoutes(server)
}
