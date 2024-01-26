package dto

const requestURL = "/api/v1/users/create-user"

const DELIVERY_ROLE = "DELIVERY"

type CreateAccountRequest struct {
	AuthorizationHeader
	Body struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		PhoneNumber string `json:"phoneNumber"`
		Email       string `json:"email"`
		Role        string `json:"role"`
	}
}

type CreateAccountResponse struct {
	Id             string `json:"id"`
	PhoneNumber    string `json:"phoneNumber"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
	Role           string `json:"role"`
}

func (CreateAccountRequest) URL() string {
	return requestURL
}
