package response

import "go-to-chat/app/user"

type MatchUserResponse struct {
	MatchedUser *user.UserResponse `json:"matched_user"`
}
