package contract

import (
	"github.com/smhdhsn/restaurant-gateway/internal/service/dto"
)

// EdibleMenuService is the interface that edible's menu service must implement.
type EdibleMenuService interface {
	List() ([]*dto.Menu, error)
}
