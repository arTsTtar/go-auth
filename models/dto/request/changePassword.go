package request

import (
	"errors"
)

type ChangePassword struct {
	Password   *string `json:"password"`
	Disable2FA bool    `json:"disable2fa"`
}

func ValidateChangePassword(changePassword ChangePassword) error {
	if changePassword.Password == nil {
		return errors.New("new password not provided")
	}
	if len(*changePassword.Password) <= 6 {
		return errors.New("user password must be at least 7 symbols long and can't be empty")
	}
	return nil
}
