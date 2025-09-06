package handlers

type ErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}

type BaseErrorResponse struct {
	Error string `json:"error"`
}
