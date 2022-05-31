package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// UserSourceRepo is the interface representing user source's repository or it's mock.
type UserSourceRepository interface {
	Store(*model.UserDTO) (*model.UserDTO, error)
	Find(*model.UserDTO) (*model.UserDTO, error)
	Destroy(*model.UserDTO) error
	Update(*model.UserDTO) error
}
