package contract

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
)

// This section holds errors that might happen within repository layer.
var (
	ErrLackOfComponents = errors.New("lack_of_components")
)

// OrderSubmitService is the interface that order's submission service must implement.
type OrderSubmitService interface {
	Submit(*dto.Order) error
}
