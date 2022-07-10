package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
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
	mListEntity, err := s.menuServ.List()
	if err != nil {
		log.Error(err)
		resp := response.NewStatusInternalServerError(nil)
		c.JSON(resp.Status, resp)
		return
	}

	mListResp := multipleMenuEntityToResp(mListEntity)

	resp := response.NewStatusOK(mListResp)
	c.JSON(resp.Status, resp)
}

// multipleMenuEntityToResp is responsible for transforming a list of menu entity to menu list response struct.
func multipleMenuEntityToResp(mListEntity []*entity.Menu) []*MenuListResp {
	mListResp := make([]*MenuListResp, len(mListEntity))

	for i, mEntity := range mListEntity {
		iListResp := make([]string, len(mEntity.Ingredients))

		copy(iListResp, mEntity.Ingredients)

		mListResp[i] = &MenuListResp{
			FoodID:      mEntity.ID,
			Title:       mEntity.Title,
			Ingredients: iListResp,
		}
	}

	return mListResp
}
