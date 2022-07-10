package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
)

// EdibleInventoryService is the interface that edible's inventory service must implement.
type EdibleInventoryService interface {
	Recycle(*dto.Recycle) error
	Buy(*dto.Buy) error
}
