package response

type UserCreationResponse struct {
	Id             uint      `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	TwoFactEnabled bool      `json:"twoFaEnabled"`
	Secret         string    `json:"secret,omitempty"`
	QrCode         string    `json:"qrCode,omitempty"`
	BackupCodes    [6]string `json:"backupCodes"`
}
