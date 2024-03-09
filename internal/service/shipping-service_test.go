package service

import (
	"better-shipping-app/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"testing"
)

type packServiceMock struct {
	mock.Mock
}

func (m *packServiceMock) ListPacks() ([]model.Pack, error) {
	args := m.Called()
	return args.Get(0).([]model.Pack), args.Error(1)
}

func checkShippingVariant(service ShippingService, t *testing.T, givenItemsCount int, expectedShippingStr string) {
	t.Run("Item Count: "+strconv.Itoa(givenItemsCount), func(t *testing.T) {
		shipping, err := service.CalculateShipping(givenItemsCount)

		if err != nil {
			t.Error(err)
		}

		if !assert.EqualValues(t, expectedShippingStr, shipping.String()) {
			t.Errorf("Expected %s, got %s", expectedShippingStr, shipping.String())
		}
	})
}

func TestShippingUnitsPacksVariant1(t *testing.T) {
	mockPackService := &packServiceMock{}

	mockPackService.On("ListPacks").Return([]model.Pack{
		{ItemCount: 250},
		{ItemCount: 500},
		{ItemCount: 1000},
		{ItemCount: 2000},
		{ItemCount: 5000},
	}, nil)

	service := NewShippingService(mockPackService)

	checkShippingVariant(service, t, 1, "1x250")
	checkShippingVariant(service, t, 250, "1x250")
	checkShippingVariant(service, t, 251, "1x500")
	checkShippingVariant(service, t, 501, "1x500 1x250")
	checkShippingVariant(service, t, 12001, "2x5000 1x2000 1x250")
	checkShippingVariant(service, t, 12011, "2x5000 1x2000 1x250")
	checkShippingVariant(service, t, 12249, "2x5000 1x2000 1x250")
	checkShippingVariant(service, t, 13490, "2x5000 1x2000 1x1000 1x500")

	mockPackService.AssertExpectations(t)
}

func TestShippingUnitsPacksVariant2(t *testing.T) {
	mockPackService := &packServiceMock{}

	mockPackService.On("ListPacks").Return([]model.Pack{
		{ItemCount: 2},
		{ItemCount: 4},
		{ItemCount: 7},
		{ItemCount: 9},
		{ItemCount: 12},
	}, nil)

	service := NewShippingService(mockPackService)

	checkShippingVariant(service, t, 1, "1x2")
	checkShippingVariant(service, t, 2, "1x2")
	checkShippingVariant(service, t, 3, "1x4")
	checkShippingVariant(service, t, 4, "1x4")
	checkShippingVariant(service, t, 11, "1x9 1x2")
	checkShippingVariant(service, t, 17, "1x12 1x4 1x2")
	checkShippingVariant(service, t, 21, "1x12 1x9")
	checkShippingVariant(service, t, 27, "2x12 1x4")
	checkShippingVariant(service, t, 29, "2x12 1x4 1x2")
	checkShippingVariant(service, t, 31, "2x12 1x7")
	checkShippingVariant(service, t, 37, "3x12 1x2")
	checkShippingVariant(service, t, 39, "3x12 1x4")
	checkShippingVariant(service, t, 37218738737182783, "3101561561431898x12 1x7")

	mockPackService.AssertExpectations(t)
}

func TestShippingWithZeroItemCount(t *testing.T) {
	mockPackService := &packServiceMock{}

	service := NewShippingService(mockPackService)

	checkShippingVariant(service, t, 0, "")

	mockPackService.AssertExpectations(t)
}

func TestShippingWithNegativeItemCount(t *testing.T) {
	mockPackService := &packServiceMock{}

	service := NewShippingService(mockPackService)

	_, err := service.CalculateShipping(-1)

	if err == nil {
		t.Error("Expected error")
	}

	if err.Error() != "item count must be a positive integer" {
		t.Errorf("Expected `item count must be a positive integer`, got `%s`", err.Error())
	}

	mockPackService.AssertExpectations(t)
}

func TestShipmentShouldFailWhenPackServiceFailsToReturnPackList(t *testing.T) {
	mockPackService := &packServiceMock{}

	mockPackService.On("ListPacks").Return([]model.Pack{}, assert.AnError)

	service := NewShippingService(mockPackService)

	_, err := service.CalculateShipping(1)

	if err == nil {
		t.Error("Expected error")
	}

	mockPackService.AssertExpectations(t)
}
