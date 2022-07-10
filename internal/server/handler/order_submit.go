package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/smhdhsn/restaurant-gateway/internal/server/helper"
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
	"github.com/smhdhsn/restaurant-gateway/pkg/response"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// This section holds errors that might happen within this handler.
var (
	ErrInvalidRouteParam = errors.New("invalid_route_parameter")
)

// OrderSubmissionHandler holds the services that will be used within this handler.
type OrderSubmissionHandler struct {
	submitServ serviceContract.OrderSubmitService
}

// NewOrderSubmissionHandler creates an instance of OrderSubmissionHandler with its dependencies.
func NewOrderSubmissionHandler(os serviceContract.OrderSubmitService) *OrderSubmissionHandler {
	return &OrderSubmissionHandler{
		submitServ: os,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *OrderSubmissionHandler) Submit(c *gin.Context) {
	foodID, err := helper.StrToUint(c.Params.ByName("foodID"))
	if err != nil {
		resp := response.NewStatusBadRequest(nil, ErrInvalidRouteParam)
		c.JSON(resp.Status, resp)
		return
	}

	oDTO := &dto.Order{
		FoodID: foodID,
		UserID: 1, // TODO: use user's id from auth middleware.
	}

	err = s.submitServ.Submit(oDTO)
	if err != nil {
		if errors.Is(err, serviceContract.ErrLackOfComponents) {
			resp := response.NewStatusNotFound(err)
			c.JSON(resp.Status, resp)
			return
		}

		log.Error(err)
		resp := response.NewStatusInternalServerError(nil)
		c.JSON(resp.Status, resp)
		return
	}

	c.Status(http.StatusNoContent)
}
