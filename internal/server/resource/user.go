package resource

import (
	uHand "github.com/smhdhsn/restaurant-gateway/internal/server/handler/user"
)

// UserResource holds handlers for user resource.
type UserResource struct {
	SourHand *uHand.UserSourceHandler
}

// NewUserResource creates an instance of UserSource with its dependencies.
func NewUserResource(sourHand *uHand.UserSourceHandler) *UserResource {
	return &UserResource{
		SourHand: sourHand,
	}
}
