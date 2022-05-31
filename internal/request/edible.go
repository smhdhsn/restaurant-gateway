package request

// EdibleRecipeReq holds the schema for edible's recipe service.
type EdibleRecipeReq struct {
	Foods []struct {
		Title       string   `json:"title" validate:"required"`
		Ingredients []string `json:"ingredients" validate:"required"`
	} `json:"foods" validate:"required"`
}
