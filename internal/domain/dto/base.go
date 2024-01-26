package dto

type AuthorizationHeader struct {
	BearerToken string `reqHeader:"Authorization"`
}

type DefaultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
