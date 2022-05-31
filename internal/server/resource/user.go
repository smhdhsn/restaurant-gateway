package resource

import (
	"github.com/smhdhsn/restaurant-gateway/internal/server/handler"
)

// UserResource holds handlers for user resource.
type UserResource struct {
	SourceHand *handler.UserSourceHandler
}

// NewUserResource creates an instance of UserSource with its dependencies.
func NewUserResource(sh *handler.UserSourceHandler) *UserResource {
	return &UserResource{
		SourceHand: sh,
	}
}
