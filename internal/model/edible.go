package model

// MenuItemDTO represents menu item's data transfer object.
type MenuItemDTO struct {
	ID                  uint32
	Title               string
	IngredientTitleList []string
}

// menuResp is the response schema.
type menuResp struct {
	ID          uint32   `json:"food_id"`
	Title       string   `json:"food_title"`
	Ingredients []string `json:"ingredients"`
}

// ToResp creates a response from MenuItemDTO.
func (s *MenuItemDTO) ToResp() menuResp {
	return menuResp{
		ID:          s.ID,
		Title:       s.Title,
		Ingredients: s.IngredientTitleList,
	}
}

// MenuItemListDTO holds a list of ItemDTO.
type MenuItemListDTO []*MenuItemDTO

// ToResp creates a response from a list of MenuItemDTOs.
func (s *MenuItemListDTO) ToResp() []menuResp {
	iList := make([]menuResp, len(*s))
	for i, item := range *s {
		iList[i] = item.ToResp()
	}

	return iList
}
