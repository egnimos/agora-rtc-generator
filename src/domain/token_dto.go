package domain

type RestError struct {
	Message string `json:"message"`
	Status  int `json:"status"`
}
