package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
	"github.com/smhdhsn/restaurant-gateway/pkg/response"

	log "github.com/smhdhsn/restaurant-gateway/internal/logger"
	serviceContract "github.com/smhdhsn/restaurant-gateway/internal/service/contract"
)

// EdibleMenuHandler holds the services that will be used within this handler.
type EdibleMenuHandler struct {
	menuServ serviceContract.EdibleMenuService
}

// NewEdibleMenuHandler creates an instance of EdibleMenuHandler with its dependencies.
func NewEdibleMenuHandler(s serviceContract.EdibleMenuService) *EdibleMenuHandler {
	return &EdibleMenuHandler{
		menuServ: s,
	}
}

// MenuListResp is the response schema of the list api of the menu handler.
type MenuListResp struct {
	FoodID      uint32   `json:"food_id"`
	Title       string   `json:"food_title"`
	Ingredients []string `json:"ingredients"`
}

// List returns a list of available items to order.
func (s *EdibleMenuHandler) List(c *gin.Context) {
	mListDTO, err := s.menuServ.List()
	if err != nil {
		log.Error(err)
		resp := response.NewStatusInternalServerError(nil)
		c.JSON(resp.Status, resp)
		return
	}

	mListResp := multipleMenuDTOToResp(mListDTO)

	resp := response.NewStatusOK(mListResp)
	c.JSON(resp.Status, resp)
}

// multipleMenuDTOToResp is responsible for transforming a list of menu dto into menu list response struct.
func multipleMenuDTOToResp(mListDTO []*dto.Menu) []*MenuListResp {
	mListResp := make([]*MenuListResp, len(mListDTO))

	for i, mDTO := range mListDTO {
		iListResp := make([]string, len(mDTO.Ingredients))

		copy(iListResp, mDTO.Ingredients)

		mListResp[i] = &MenuListResp{
			FoodID:      mDTO.ID,
			Title:       mDTO.Title,
			Ingredients: iListResp,
		}
	}

	return mListResp
}
