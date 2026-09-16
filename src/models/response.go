package models

// Response is the shared JSON shape returned by any endpoint that responds with JSON.
type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
