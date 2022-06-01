package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/restaurant-gateway/internal/model"
	"github.com/smhdhsn/restaurant-gateway/internal/server/helper"
	"github.com/smhdhsn/restaurant-gateway/pkg/response"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// OrderSubmitHandler holds the services that will be used within this handler.
type OrderSubmitHandler struct {
	serv serviceContract.OrderSubmitService
}

// NewOrderSubmitHandler creates an instance of OrderSubmitHandler with its dependencies.
func NewOrderSubmitHandler(os serviceContract.OrderSubmitService) *OrderSubmitHandler {
	return &OrderSubmitHandler{
		serv: os,
	}
}

// Submit is responsible for submiting an order which includes storing order details inside database, and decreasing related component's stocks.
func (s *OrderSubmitHandler) Submit(c *gin.Context) {
	foodID, err := helper.StrToUint(c.Params.ByName("foodID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("invalid route parameter"))
		return
	}

	oReq := model.OrderDTO{
		FoodID: foodID,
		UserID: uint32(1),
	}

	err = s.serv.Submit(&oReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrUncaught) {
			log.Error(err)
			c.JSON(response.NewStatusNotImplemented())
		} else {
			log.Error(err)
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}
