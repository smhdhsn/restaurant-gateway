package entity

import (
	"time"
)

// Recycle represents repository's entity struct.
type Recycle struct {
	Expired  bool
	Finished bool
}

// Buy represents repository's entity struct.
type Buy struct {
	Amount    uint32
	ExpiresAt time.Time
}
