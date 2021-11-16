package entity

import "time"

type User struct {
	Id             uint
	Name           string       `gorm:"notnull"`
	Email          string       `gorm:"unique;notnull"`
	Password       []byte       `json:"-" gorm:"notnull"`
	TwoFactEnabled bool         `gorm:"notnull;default=false"`
	TwoFactSecret  string       `json:"-"`
	Roles          []Role       `json:"-" gorm:"many2many:user_roles;"`
	BackupCodes    []BackupCode `gorm:"foreignKey:UserId; constraint:OnUpdate:CASCADE, OnDelete:CASCADE" json:"-"`
	CreatedAt      time.Time    `json:"-"`
	UpdatedAt      time.Time    `json:"-"`
}

type Users []*User
