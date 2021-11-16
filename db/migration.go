package db

import (
	"go-auth/db/seeds"
	"go-auth/entity"
)

func Automigrate() {

	err := DB.AutoMigrate(&entity.User{}, &entity.BackupCode{}, &entity.Role{})
	if err != nil {
		panic("could not connect to db")
		return
	}
	seeds.Execute(DB)

}
