package entity

type User struct {
	Id             uint
	Name           string       `gorm:"notnull"`
	Email          string       `gorm:"unique;notnull"`
	Password       []byte       `json:"-" gorm:"notnull"`
	TwoFactEnabled bool         `gorm:"notnull;default=false"`
	TwoFactSecret  string       `json:"-"`
	BackupCodes    []BackupCode `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE" json:"-"`
}

type Users []*User
