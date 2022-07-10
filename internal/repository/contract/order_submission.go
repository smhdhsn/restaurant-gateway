package contract

import (
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
)

// This section holds errors that might happen within repository layer.
var (
	ErrLackOfComponents = errors.New("lack_of_components")
)

// OrderSubmissionRepository is the interface that order service must implement.
type OrderSubmissionRepository interface {
	Submit(*entity.Order) error
}
