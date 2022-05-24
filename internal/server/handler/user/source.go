package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/smhdhsn/restaurant-gateway/internal/model"
	"github.com/smhdhsn/restaurant-gateway/internal/server/helper"
	"github.com/smhdhsn/restaurant-gateway/util/response"

	uSorcRepoContract "github.com/smhdhsn/restaurant-gateway/internal/repository/contract/user"
	uSorcRequst "github.com/smhdhsn/restaurant-gateway/internal/request/user"
	uServContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract/user"
)

// UserSourceHandler holds the services that will be used within this handler.
type UserSourceHandler struct {
	sourceServ uServContract.UserSourceService
}

// NewUserSourceHandler creates an instance of UserSourceHandler with its dependencies.
func NewUserSourceHandler(sorc uServContract.UserSourceService) *UserSourceHandler {
	return &UserSourceHandler{
		sourceServ: sorc,
	}
}

// Store is responsible for storing a user into  database.
func (s *UserSourceHandler) Store(c *gin.Context) {
	req := new(uSorcRequst.SourceStoreReq)
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

	uDTO, err := s.sourceServ.Store(uReq)
	if err != nil {
		if errors.Is(err, uSorcRepoContract.ErrDuplicateEntry) {
			c.JSON(response.NewStatusBadRequest(err.Error()))
		} else if errors.Is(err, uSorcRepoContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
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

	uDTO, err := s.sourceServ.Find(uReq)
	if err != nil {
		if errors.Is(err, uSorcRepoContract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, uSorcRepoContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	data := uDTO.ToResp()
	c.JSON(response.NewStatusOK(data))
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

	err = s.sourceServ.Destroy(uReq)
	if err != nil {
		if errors.Is(err, uSorcRepoContract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, uSorcRepoContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// Update is responsible for updating user's information inside database.
func (s *UserSourceHandler) Update(c *gin.Context) {
	req := new(uSorcRequst.SourceUpdateReq)
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

	err = s.sourceServ.Update(uReq)
	if err != nil {
		if errors.Is(err, uSorcRepoContract.ErrRecordNotFound) {
			c.JSON(response.NewStatusNotFound())
		} else if errors.Is(err, uSorcRepoContract.ErrDuplicateEntry) {
			c.JSON(response.NewStatusBadRequest(err.Error()))
		} else if errors.Is(err, uSorcRepoContract.ErrUncaught) {
			c.JSON(response.NewStatusNotImplemented())
		} else {
			c.JSON(response.NewStatusInternalServerError())
		}
		return
	}

	c.Status(http.StatusNoContent)
}
