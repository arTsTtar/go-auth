package orm

type User struct {
	Id             uint
	Name           string `gorm:"notnull"`
	Email          string `gorm:"unique;notnull"`
	Password       []byte `gorm:"notnull"`
	TwoFactEnabled bool   `gorm:"notnull;default=false"`
	TwoFactSecret  string
}
