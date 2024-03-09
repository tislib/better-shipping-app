package service

import (
	"better-shipping-app/internal/dto"
	"better-shipping-app/internal/model"
	"cmp"
	"errors"
	"slices"
)

type ShippingService interface {
	CalculateShipping(itemCount int) (dto.Shipping, error)
}

type shippingService struct {
	packService PackService
}

func (s shippingService) CalculateShipping(itemCount int) (dto.Shipping, error) {
	// check item count
	if itemCount < 0 {
		return dto.Shipping{}, errors.New("item count must be a positive integer")
	}

	if itemCount == 0 {
		return dto.Shipping{}, nil
	}

	// get packs
	packs, err := s.packService.ListPacks()

	if err != nil {
		return dto.Shipping{}, err
	}

	// sort packs by size descending
	slices.SortFunc(packs, func(a, b model.Pack) int {
		return cmp.Compare(b.ItemCount, a.ItemCount)
	})

	// calculate refill count
	// refill count is the number of items that need to be added to the order to make it suitable for packing. (like minimum amount of items to be added to the order to make it suitable for packing)
	var refillCount = s.calculateRefillCount(itemCount, packs)

	// remainingItemCount is actual count after refill.
	remainingItemCount := itemCount + refillCount

	var shipping dto.Shipping
	for _, pack := range packs {
		// calculate how many packs are needed
		count := remainingItemCount / pack.ItemCount

		if count > 0 {
			shipping.Items = append(shipping.Items, dto.ShippingItem{
				Pack:  pack,
				Count: count,
			})
		}

		// calculate remaining item count
		remainingItemCount = remainingItemCount % pack.ItemCount

		// if the remainingItemCount is 0, then we are done
		if remainingItemCount == 0 {
			break
		}
	}

	return shipping, nil
}

// calculateRefillCount calculates the number of items that need to be added to the order to make it suitable for packing.
func (s shippingService) calculateRefillCount(itemCount int, packs []model.Pack) int {
	var remainingItemCount = itemCount

	// decrease the remainingItemCount by the pack size as much as possible
	for _, pack := range packs {
		remainingItemCount = remainingItemCount % pack.ItemCount

		if remainingItemCount == 0 {
			break
		}
	}

	// if the remainingItemCount is not 0, then we need
	// to add the remainingItemCount to the order to make it suitable for packing
	if remainingItemCount != 0 {
		var smallestPack = packs[len(packs)-1]
		return smallestPack.ItemCount - remainingItemCount
	}

	return 0
}

func NewShippingService(packService PackService) ShippingService {
	return &shippingService{packService: packService}
}
