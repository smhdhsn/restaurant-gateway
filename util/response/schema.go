package response

// dataResp is the struct holding response body containing 'data' key.
type dataResp struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}

// messageResp is the struct holding response body containing 'message' key.
type messageResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
