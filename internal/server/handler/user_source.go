package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/smhdhsn/restaurant-gateway/internal/model"
	"github.com/smhdhsn/restaurant-gateway/internal/request"
	"github.com/smhdhsn/restaurant-gateway/internal/server/helper"
	"github.com/smhdhsn/restaurant-gateway/pkg/response"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	repositoryContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// UserSourceHandler holds the services that will be used within this handler.
type UserSourceHandler struct {
	serv serviceContract.UserSourceService
}

// NewUserSourceHandler creates an instance of UserSourceHandler with its dependencies.
func NewUserSourceHandler(s serviceContract.UserSourceService) *UserSourceHandler {
	return &UserSourceHandler{
		serv: s,
	}
}

// Store is responsible for storing a user into  database.
func (s *UserSourceHandler) Store(c *gin.Context) {
	req := new(request.SourceStoreReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("invalid json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	uReq := &model.UserDTO{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Status:    req.Status,
	}

	uDTO, err := s.serv.Store(uReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrDuplicateEntry) {
			c.JSON(response.NewStatusBadRequest(err.Error()))
		} else if errors.Is(err, repositoryContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			log.Error(err)
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := uDTO.ToResp()
	c.JSON(response.NewStatusCreated(data))
}

// Find is responsible for fetching user information from database.
func (s *UserSourceHandler) Find(c *gin.Context) {
	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("invalid route parameter"))
		return
	}

	uReq := &model.UserDTO{
		ID: userID,
	}

	uDTO, err := s.serv.Find(uReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, repositoryContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			log.Error(err)
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := uDTO.ToResp()
	c.JSON(response.NewStatusOKWithData(data))
}

// Destroy is responsible for deleting a user from the database.
func (s *UserSourceHandler) Destroy(c *gin.Context) {
	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("invalid route parameter"))
		return
	}

	uReq := &model.UserDTO{
		ID: userID,
	}

	err = s.serv.Destroy(uReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, repositoryContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			log.Error(err)
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// Update is responsible for updating user's information inside database.
func (s *UserSourceHandler) Update(c *gin.Context) {
	req := new(request.SourceUpdateReq)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(response.NewStatusBadRequest("invalid json body"))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		data := helper.ParseValidationErr(err)
		c.JSON(response.NewStatusUnprocessableEntity(data))
		return
	}

	userID, err := helper.StrToUint(c.Params.ByName("userID"))
	if err != nil {
		c.JSON(response.NewStatusBadRequest("invalid route parameter"))
		return
	}

	uReq := &model.UserDTO{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Status:    req.Status,
	}

	err = s.serv.Update(uReq)
	if err != nil {
		if errors.Is(err, repositoryContract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, repositoryContract.ErrDuplicateEntry) {
			c.JSON(response.NewStatusBadRequest(err.Error()))
		} else if errors.Is(err, repositoryContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			log.Error(err)
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}
