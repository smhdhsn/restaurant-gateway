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
	uRepoContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract/user"
)

// UserRepo contains repository's database connection.
type UserSourceRepo struct {
	usc uspb.UserSourceServiceClient
	ctx *context.Context
}

// NewUserSourceRepository creates an instance of the remote repository with gRPC connection.
func NewUserSourceRepository(ctx *context.Context, conn uspb.UserSourceServiceClient) uRepoContract.UserSourceRepository {
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
				return uRepoContract.ErrRecordNotFound
			case codes.AlreadyExists:
				return uRepoContract.ErrDuplicateEntry
			case codes.Unknown:
				return uRepoContract.ErrUncaught
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
				return nil, uRepoContract.ErrDuplicateEntry
			case codes.Unknown:
				return nil, uRepoContract.ErrUncaught
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
				return nil, uRepoContract.ErrRecordNotFound
			case codes.Unknown:
				return nil, uRepoContract.ErrUncaught
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
				return uRepoContract.ErrRecordNotFound
			case codes.Unknown:
				return uRepoContract.ErrUncaught
			}
		}

		return errors.Wrap(err, "error on calling destroy on user gRPC server")
	}

	return nil
}
