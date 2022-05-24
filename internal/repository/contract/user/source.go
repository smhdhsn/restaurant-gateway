package user

import (
	"errors"

	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// This block holds common errors that might happen within user source repository.
var (
	ErrUncaught       = errors.New("uncought error")
	ErrRecordNotFound = errors.New("record not found")
	ErrDuplicateEntry = errors.New("duplicate entry")
)

// UserSourceRepo is the interface representing user source's repository or it's mock.
type UserSourceRepository interface {
	Store(*model.UserDTO) (*model.UserDTO, error)
	Find(*model.UserDTO) (*model.UserDTO, error)
	Destroy(*model.UserDTO) error
	Update(*model.UserDTO) error
}
