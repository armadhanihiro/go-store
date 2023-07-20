package schema

type SignInReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenReq struct {
	UserID       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResp struct {
	AccessToken string `json:"access_token"`
}
