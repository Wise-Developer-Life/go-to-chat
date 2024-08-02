package user

import "go-to-chat/app/model"

type UserResponse struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfileUrl string `json:"profile_url"`
}

func NewUserResponse(user *model.User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		ProfileUrl: user.ProfileUrl,
	}
}
