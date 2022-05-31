package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// EdibleMenuService is the interface that edible's menu service must implement.
type EdibleMenuService interface {
	List() (model.MenuItemListDTO, error)
}
