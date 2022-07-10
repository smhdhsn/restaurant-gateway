package dto

import (
	"time"
)

// Recycle represents recycle's data transfer object.
type Recycle struct {
	Expired  bool
	Finished bool
}

// Buy represents buy's data transfer object.
type Buy struct {
	Amount    uint32
	ExpiresAt time.Time
}
