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

	minPackSize := packs[len(packs)-1].ItemCount

	// identify top as the maximum possible item count,
	// which can be the case if the itemCount is packed only with the smallest pack
	top := (itemCount/minPackSize + 1) * minPackSize
	// packMap is a map of pack size to pack
	var packMap = make(map[int]model.Pack)

	// dp is a slice of size top+1, which will be used to store the minimum number of packs required to pack the itemCount
	dp := make([]int, top+1)

	result := make([]int, top+1)

	// initialize dp with top+1
	for i := range dp {
		dp[i] = top + 1
	}

	dp[0] = 0

	// the purpose here is to find the minimum number of packs required to pack the itemCount, and store the result in dp
	// it will continue to calculate even after itemCount reaches, until top.
	for _, pack := range packs {
		var packSize = pack.ItemCount
		packMap[packSize] = pack

		for i := packSize; i <= top; i++ {
			// if pack count is less than the current minimum pack count, update
			if dp[i-packSize]+1 < dp[i] && dp[i-packSize] != top+1 {
				dp[i] = dp[i-packSize] + 1
				result[i] = packSize
			}
		}
	}

	// locate the nearest found itemCount
	nearest := itemCount
	for i := itemCount; i < top+1; i++ {
		if dp[i] != top+1 {
			nearest = i
			break
		}
	}

	packCountMap := make(map[int]int, len(packs))

	for i := nearest; i > 0; i -= result[i] {
		packCountMap[result[i]]++
	}

	shipping := dto.Shipping{}

	for packSize, count := range packCountMap {
		shipping.Items = append(shipping.Items, dto.ShippingItem{
			Pack:  packMap[packSize],
			Count: count,
		})

	}

	return shipping, nil
}

func NewShippingService(packService PackService) ShippingService {
	return &shippingService{packService: packService}
}
