package response

type PasswordResetResponse struct {
	Email             *string `json:"email"`
	Name              *string `json:"name"`
	GeneratedPassword *string `json:"generatedPassword"`
}

func ToPasswordResetResponse(email *string, generatedPassword *string, name *string) (passwordResetResponse PasswordResetResponse) {
	return PasswordResetResponse{
		Email:             email,
		GeneratedPassword: generatedPassword,
		Name:              name,
	}
}
