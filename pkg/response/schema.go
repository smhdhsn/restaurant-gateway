package response

// Schema is the response schema for returning data to client.
type Schema struct {
	Status int     `json:"-"`
	Error  *string `json:"error"`
	Data   any     `json:"data"`
}
