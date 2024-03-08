package service

type ShippingService interface {
}

type shippingService struct {
	packService PackService
}

func NewShippingService(packService PackService) ShippingService {
	return &shippingService{packService: packService}
}
