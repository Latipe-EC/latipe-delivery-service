package dto

type AuthorizationHeader struct {
	BearerToken string `reqHeader:"Authorization"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type BaseHeader struct {
	BearToken string `json:"bear_token"`
}
