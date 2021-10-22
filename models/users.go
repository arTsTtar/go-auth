package models

type User struct {
	Id       uint
	Name     string `gorm:"notnull"`
	Email    string `gorm:"unique;notnull"`
	Password []byte `gorm:"notnull"`
	Secret   []byte `gorm:"notnull"`
}
