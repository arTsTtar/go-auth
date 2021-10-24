package repository

import (
	"go-auth/models/entity"
	"gorm.io/gorm"
	"log"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	FindUserByEmail(email string) (entity.User, error)
	FindUserById(id string) (entity.User, error)
	Save(user entity.User) (entity.User, error)
	FindAllUsers() (entity.Users, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (u userRepository) FindUserByEmail(email string) (user entity.User, err error) {
	log.Print("[UserRepository]...FindUserByEmail")
	err = u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u userRepository) FindAllUsers() (users entity.Users, err error) {
	log.Print("[UserRepository]...FindAllUsers")
	err = u.DB.Find(&users).Error
	return users, err
}

func (u userRepository) FindUserById(id string) (user entity.User, err error) {
	log.Print("[UserRepository]...FindUserById")
	err = u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func (u userRepository) Save(user entity.User) (entity.User, error) {
	log.Print("[UserRepository]...Save")
	err := u.DB.Create(&user).Error
	return user, err
}
