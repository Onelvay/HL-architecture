package dto

type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignUpResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	AccessToken string `json:"accessToken"`
	//refresh_token
}
