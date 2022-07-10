package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
)

// EdibleMenuRepository is the interface representing edible menu's repository or it's mock.
type EdibleMenuRepository interface {
	List() ([]*entity.Menu, error)
}
