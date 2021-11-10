package response

import "go-auth/models/entity"

type SimpleUserResponse struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	TwoFactEnabled bool   `json:"twoFaEnabled"`
}

func ToSimpleUserResponse(user *entity.User) (userResponse SimpleUserResponse) {
	return SimpleUserResponse{
		Id:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		TwoFactEnabled: user.TwoFactEnabled,
	}
}
