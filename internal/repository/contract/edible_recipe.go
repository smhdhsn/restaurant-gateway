package contract

import "github.com/smhdhsn/restaurant-gateway/internal/model"

// EdibleRecipeRepository is the interface representing edible menu's repository or it's mock.
type EdibleRecipeRepository interface {
	Store(model.MenuItemListDTO) error
}
