package user

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"

	uRepoContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract/user"
	uServContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract/user"
)

// UserSourceServ contains repositories that will be used within this service.
type UserSourceServ struct {
	usr uRepoContract.UserSourceRepository
}

// NewUserSourceService creates a user's source service with it's dependencies.
func NewUserSourceService(usr uRepoContract.UserSourceRepository) uServContract.UserSourceService {
	return &UserSourceServ{
		usr: usr,
	}
}

// Store is responsible for calling Store API on user source gRPC server.
func (s *UserSourceServ) Store(u *model.UserDTO) (*model.UserDTO, error) {
	return s.usr.Store(u)
}

// Find is responsible for calling Find API on user source gRPC server.
func (s *UserSourceServ) Find(u *model.UserDTO) (*model.UserDTO, error) {
	return s.usr.Find(u)
}

// Destroy is responsible for calling Destroy API on user source gRPC server.
func (s *UserSourceServ) Destroy(u *model.UserDTO) error {
	return s.usr.Destroy(u)
}

// Update is responsible for calling Update API on user source gRPC server.
func (s *UserSourceServ) Update(u *model.UserDTO) error {
	return s.usr.Update(u)
}
