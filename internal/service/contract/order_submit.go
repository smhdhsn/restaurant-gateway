package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// OrderSubmitService is the interface that order's submission service must implement.
type OrderSubmitService interface {
	Submit(*model.OrderDTO) error
}
