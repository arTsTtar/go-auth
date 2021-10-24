package request

type UserRequest struct {
	Name           *string `json:"name"`
	Email          *string `json:"email"`
	Password       *string `json:"password"`
	TwoFactEnabled bool    `json:"twoFaEnabled"`
}

func ValidateCreation(userRequest UserRequest) bool {
	if userRequest.Name == nil || len(*userRequest.Name) == 0 {
		return false
	}
	if userRequest.Email == nil || len(*userRequest.Email) < 6 {
		return false
	}
	if userRequest.Password == nil || len(*userRequest.Password) < 8 {
		return false
	}
	return true
}
