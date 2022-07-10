package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/repository/entity"
)

// EdibleMenuService is the interface that edible's menu service must implement.
type EdibleMenuService interface {
	List() ([]*entity.Menu, error)
}
