package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// OrderSubmissionRepository is the interface that order service must implement.
type OrderSubmissionRepository interface {
	Submit(o *model.OrderDTO) error
}
