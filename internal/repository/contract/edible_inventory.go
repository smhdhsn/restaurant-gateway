package contract

import (
	"time"
)

// EdibleInventoryRepository is the interface representing edible menu's repository or it's mock.
type EdibleInventoryRepository interface {
	Buy(uint32, time.Time) error
	Recycle(bool, bool) error
	Use(uint32) error
}
