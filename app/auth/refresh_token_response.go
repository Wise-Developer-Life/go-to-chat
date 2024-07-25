package auth

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewRefreshTokenResponse(accessToken string) *RefreshTokenResponse {
	return &RefreshTokenResponse{
		AccessToken: accessToken,
	}
}
