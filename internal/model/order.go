package model

import (
	"time"
)

// OrderDTO represents order's data transfer object.
type OrderDTO struct {
	ID        uint32
	UserID    uint32
	FoodID    uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OrderListDTO represents a list of OrderDTOs.
type OrderListDTO []*OrderDTO
