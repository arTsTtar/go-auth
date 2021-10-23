package orm

type User struct {
	Id             uint   `json:"id"`
	Name           string `json:"name" gorm:"notnull"`
	Email          string `json:"email" gorm:"unique;notnull"`
	Password       []byte `json:"password" gorm:"notnull"`
	TwoFactEnabled bool   `json:"twoFaEnabled" gorm:"notnull;default=false"`
	TwoFactSecret  string `json:"-"`
}
