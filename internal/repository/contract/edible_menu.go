package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// EdibleMenuRepository is the interface representing edible menu's repository or it's mock.
type EdibleMenuRepository interface {
	List() (model.MenuItemListDTO, error)
}
