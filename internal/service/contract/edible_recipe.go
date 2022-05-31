package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/model"
)

// EdibleRecipeService is the interface that edible's recipe service must implement.
type EdibleRecipeService interface {
	Store(model.MenuItemListDTO) error
}
