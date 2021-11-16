package entity

type Role struct {
	Id   uint   `gorm:"primarykey"`
	Name string `gorm:"notnull"`
}

type Roles []*Role
