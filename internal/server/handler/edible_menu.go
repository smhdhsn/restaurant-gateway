package handler

import (
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/restaurant-gateway/pkg/response"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// EdibleMenuHandler holds the services that will be used within this handler.
type EdibleMenuHandler struct {
	serv serviceContract.EdibleMenuService
}

// NewEdibleMenuHandler creates an instance of EdibleMenuHandler with its dependencies.
func NewEdibleMenuHandler(s serviceContract.EdibleMenuService) *EdibleMenuHandler {
	return &EdibleMenuHandler{
		serv: s,
	}
}

// List returns a list of available items to order.
func (s *EdibleMenuHandler) List(c *gin.Context) {
	iListDTO, err := s.serv.List()
	if err != nil {
		if errors.Is(err, repositoryContract.ErrEmptyResult) {
			c.JSON(response.NewStatusOKWithMessage(err.Error()))
		} else if errors.Is(err, repositoryContract.ErrUncaught) {
			log.Error(err)
			c.JSON(response.NewStatusNotImplemented())
		} else {
			log.Error(err)
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := iListDTO.ToResp()
	c.JSON(response.NewStatusOKWithData(data))
}
