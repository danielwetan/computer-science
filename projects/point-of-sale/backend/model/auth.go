package model

type LoginRequest struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"min=8,required"`
}

type LoginResponse struct {
	JWTToken string `json:"jwt_token"`
}
