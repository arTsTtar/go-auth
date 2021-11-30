package response

import (
	"go-auth/entity"
)

type SimpleUserResponse struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	TwoFactEnabled bool   `json:"twoFaEnabled"`
	UpdatedAt      string `json:"updatedAt"`
}

func ToSimpleUserResponse(user *entity.User) (userResponse SimpleUserResponse) {
	return SimpleUserResponse{
		Id:             user.Id,
		Name:           user.Name,
		Email:          user.Email,
		TwoFactEnabled: user.TwoFactEnabled,
		UpdatedAt:      user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

type SimpleUserResponseArray []SimpleUserResponse
