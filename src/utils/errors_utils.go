package utils

type ApplicationError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
