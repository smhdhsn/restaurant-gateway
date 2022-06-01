package resource

import (
	"github.com/smhdhsn/restaurant-gateway/internal/server/handler"
)

// OrderResource holds handlers for order resource.
type OrderResource struct {
	OrderSubmit *handler.OrderSubmitHandler
}

// NewOrderResource creates an instance of OrderResource with its dependencies.
func NewOrderResource(os *handler.OrderSubmitHandler) *OrderResource {
	return &OrderResource{
		OrderSubmit: os,
	}
}
