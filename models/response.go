package models

type Response struct {
	Data  any           `json:"data"`
	Error ErrorResponse `json:"error"`
}
