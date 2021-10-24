package db

import (
	"go-auth/models/entity"
	"go-auth/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dbUser := utils.GoDotEnvVariable("dbUser")
	dbPass := utils.GoDotEnvVariable("MYSQL_ROOT_PASSWORD")
	dbHost := utils.GoDotEnvVariable("dbHost")
	dbPort := utils.GoDotEnvVariable("dbPort")
	dbName := utils.GoDotEnvVariable("MYSQL_DATABASE")
	connString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	connection, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the DB.")
	}

	DB = connection

	connection.AutoMigrate(&entity.User{}, &entity.BackupCode{})

}
