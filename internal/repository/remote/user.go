package remote

import (
	"context"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/model"

	uspb "github.com/smhdhsn/restaurant-gateway/internal/protos/user/source"
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

// UserRepo contains repository's database connection.
type UserSourceRepo struct {
	usc uspb.UserSourceServiceClient
	ctx *context.Context
}

// NewUserSourceRepository creates an instance of the remote repository with gRPC connection.
func NewUserSourceRepository(ctx *context.Context, conn uspb.UserSourceServiceClient) UserSourceRepository {
	return &UserSourceRepo{
		usc: conn,
		ctx: ctx,
	}
}

// Update is responsible for calling Update API on user source gRPC server.
func (s *UserSourceRepo) Update(u *model.UserDTO) error {
	req := &uspb.UserUpdateRequest{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Status:    uspb.Status(uspb.Status_value[strings.ToUpper(u.Status)]),
	}

	_, err := s.usc.Update(*s.ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				return ErrRecordNotFound
			case codes.AlreadyExists:
				return ErrDuplicateEntry
			case codes.Unknown:
				return ErrUncaught
			}
		}

		return errors.Wrap(err, "error on calling update on user gRPC server")
	}

	return nil
}

// Store is responsible for calling Store API on user source gRPC server.
func (s *UserSourceRepo) Store(u *model.UserDTO) (*model.UserDTO, error) {
	req := &uspb.UserStoreRequest{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Status:    uspb.Status(uspb.Status_value[strings.ToUpper(u.Status)]),
	}

	resp, err := s.usc.Store(*s.ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.AlreadyExists:
				return nil, ErrDuplicateEntry
			case codes.Unknown:
				return nil, ErrUncaught
			}
		}

		return nil, errors.Wrap(err, "error on calling store on user gRPC server")
	}

	uDTO := &model.UserDTO{
		ID:        resp.GetId(),
		FirstName: resp.GetFirstName(),
		LastName:  resp.GetLastName(),
		Email:     resp.GetEmail(),
		Status:    resp.GetStatus().String(),
		CreatedAt: time.Unix(resp.GetCreatedAt(), 0),
		UpdatedAt: time.Unix(resp.GetUpdatedAt(), 0),
	}

	return uDTO, nil
}

// Find is responsible for calling Find API on user source gRPC server.
func (s *UserSourceRepo) Find(u *model.UserDTO) (*model.UserDTO, error) {
	req := &uspb.UserFindRequest{
		Id: u.ID,
	}

	resp, err := s.usc.Find(*s.ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				return nil, ErrRecordNotFound
			case codes.Unknown:
				return nil, ErrUncaught
			}
		}

		return nil, errors.Wrap(err, "error on calling find on user gRPC server")
	}

	uDTO := &model.UserDTO{
		ID:        resp.GetId(),
		FirstName: resp.GetFirstName(),
		LastName:  resp.GetLastName(),
		Email:     resp.GetEmail(),
		Status:    resp.GetStatus().String(),
		CreatedAt: time.Unix(resp.GetCreatedAt(), 0),
		UpdatedAt: time.Unix(resp.GetUpdatedAt(), 0),
	}

	return uDTO, nil
}

// Destroy is responsible for calling Destroy API on user source gRPC server.
func (s *UserSourceRepo) Destroy(u *model.UserDTO) error {
	req := &uspb.UserDestroyRequest{
		Id: u.ID,
	}

	_, err := s.usc.Destroy(*s.ctx, req)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				return ErrRecordNotFound
			case codes.Unknown:
				return ErrUncaught
			}
		}

		return errors.Wrap(err, "error on calling destroy on user gRPC server")
	}

	return nil
}
