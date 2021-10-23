package response

type SimpleUserResponse struct {
	Id             uint   `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	TwoFactEnabled bool   `json:"twoFaEnabled"`
}
