package contract

import (
	"time"
)

// EdibleInventoryService is the interface that edible's inventory service must implement.
type EdibleInventoryService interface {
	Buy(uint32, time.Time) error
	Recycle(bool, bool) error
	Use(uint32) error
}
