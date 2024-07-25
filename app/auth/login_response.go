package auth

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewLoginResponse(accessToken string, refreshToken string) *LoginResponse {
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
