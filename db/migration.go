package db

import (
	"github.com/harranali/authority"
	"go-auth/models/entity"
)

func Automigrate() {

	err := DB.AutoMigrate(&entity.User{}, &entity.BackupCode{})
	if err != nil {
		panic("could not connect to db")
		return
	}

	//Initialize authority tables
	_ = authority.New(authority.Options{
		TablesPrefix: "authority_",
		DB:           DB,
	})
}
