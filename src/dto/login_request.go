package dto

type LoginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}
