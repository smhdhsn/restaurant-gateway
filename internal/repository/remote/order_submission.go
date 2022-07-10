package remote

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"

	submissionProto "github.com/smhdhsn/restaurant-gateway/internal/protos/order/submission"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
)

// OrderSubmissionRepo contains repository's database connection.
type OrderSubmissionRepo struct {
	client submissionProto.OrderSubmissionServiceClient
	ctx    context.Context
}

// NewOrderSubmissionRepository creates an instance of the remote repository with gRPC connection.
func NewOrderSubmissionRepository(ctx context.Context, conn submissionProto.OrderSubmissionServiceClient) repositoryContract.OrderSubmissionRepository {
	return &OrderSubmissionRepo{
		client: conn,
		ctx:    ctx,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (r *OrderSubmissionRepo) Submit(oEntity *entity.Order) error {
	req := singleOrderEntityToSubmitReq(oEntity)

	_, err := r.client.Submit(r.ctx, req)
	if err != nil {
		if status, ok := status.FromError(err); ok {
			switch status.Code() {
			case codes.Internal:
				return errors.Wrap(err, "error inside edible gRPC server")
			case codes.NotFound:
				return repositoryContract.ErrLackOfComponents
			}
		}

		return errors.Wrap(err, "error on calling destroy on user gRPC server")
	}

	return nil
}

// singleOrderEntityToSubmitReq is responsible for transforming an order entity into an order submit request struct.
func singleOrderEntityToSubmitReq(oEntity *entity.Order) *submissionProto.OrderSubmitRequest {
	return &submissionProto.OrderSubmitRequest{
		FoodId: oEntity.FoodID,
		UserId: oEntity.UserID,
	}
}
