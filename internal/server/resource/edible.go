package resource

import (
	"github.com/smhdhsn/restaurant-gateway/internal/server/handler"
)

// EdibleResource holds handlers for edible resource.
type EdibleResource struct {
	MenuHand *handler.EdibleMenuHandler
}

// NewEdibleResource creates an instance of UserSource with its dependencies.
func NewEdibleResource(mh *handler.EdibleMenuHandler) *EdibleResource {
	return &EdibleResource{
		MenuHand: mh,
	}
}
