package api

import (
	"better-shipping-app/internal/dto"
	"better-shipping-app/internal/service"
	"encoding/json"
	"net/http"
)

type CalculateShippingRequest struct {
	ItemCount int `json:"ItemCount"`
}

type CalculateShippingResponse struct {
	Shipping dto.Shipping `json:"shipping"`
	Text     string       `json:"text"`
}

type shippingApi struct {
	service service.ShippingService
}

func (p shippingApi) registerRoutes(s Server) {
	s.getMux().HandleFunc("POST /shipping", p.calculate)
}

func (p shippingApi) calculate(writer http.ResponseWriter, request *http.Request) {
	var requestBody = CalculateShippingRequest{}

	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		handleHttpError(writer, 400, err.Error())
		return
	}

	shipping, err := p.service.CalculateShipping(requestBody.ItemCount)

	if err != nil {
		handleHttpError(writer, 400, err.Error())
		return
	}

	var responseBody = CalculateShippingResponse{
		Shipping: shipping,
		Text:     shipping.String(),
	}

	respondJsonBody(writer, responseBody)
}

func RegisterShippingApi(shippingService service.ShippingService, server Server) {
	p := &shippingApi{service: shippingService}

	p.registerRoutes(server)
}
