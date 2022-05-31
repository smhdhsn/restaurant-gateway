package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// UserSourceService is the interface that user service must implement.
type UserSourceService interface {
	Store(*model.UserDTO) (*model.UserDTO, error)
	Find(*model.UserDTO) (*model.UserDTO, error)
	Destroy(*model.UserDTO) error
	Update(*model.UserDTO) error
}
