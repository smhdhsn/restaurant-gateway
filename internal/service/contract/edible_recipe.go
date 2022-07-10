package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
)

// EdibleRecipeService is the interface that edible's recipe service must implement.
type EdibleRecipeService interface {
	Store([]*dto.Recipe) error
}
