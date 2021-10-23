package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name" gorm:"notnull"`
	Email    string `json:"email" gorm:"unique;notnull"`
	Password []byte `json:"-" gorm:"notnull"`
	Secret   []byte `json:"-" gorm:"notnull"`
}
