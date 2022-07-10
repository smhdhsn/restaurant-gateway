package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
)

// EdibleInventoryRepository is the interface representing edible menu's repository or it's mock.
type EdibleInventoryRepository interface {
	Recycle(*entity.Recycle) error
	Buy(*entity.Buy) error
}
