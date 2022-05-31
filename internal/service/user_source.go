package service

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"

	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// UserSourceServ contains repositories that will be used within this service.
type UserSourceServ struct {
	usr repositoryContract.UserSourceRepository
}

// NewUserSourceService creates a user's source service with it's dependencies.
func NewUserSourceService(usr repositoryContract.UserSourceRepository) serviceContract.UserSourceService {
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
