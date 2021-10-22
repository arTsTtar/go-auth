package response

type UserResponse struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Secret string `json:"secret"`
	QrCode string `json:"qrCode"`
}
