package dto

import (
	"better-shipping-app/internal/model"
	"slices"
	"strconv"
	"strings"
)

type ShippingItem struct {
	Pack  model.Pack `json:"pack"`
	Count int        `json:"count"`
}

type Shipping struct {
	Items []ShippingItem `json:"items"`
}

func (s Shipping) String() string {
	var buffer []string

	slices.SortFunc(s.Items, func(a, b ShippingItem) int {
		return b.Pack.ItemCount - a.Pack.ItemCount
	})

	for _, item := range s.Items {
		buffer = append(buffer, strconv.Itoa(item.Count)+"x"+strconv.Itoa(item.Pack.ItemCount))
	}

	return strings.Join(buffer, " ")
}
