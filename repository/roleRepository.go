package repository

import (
	"go-auth/entity"
	"gorm.io/gorm"
	"log"
)

type roleRepository struct {
	DB *gorm.DB
}

type RoleRepository interface {
	FindByName(name string) (entity.Role, error)
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		DB: db,
	}
}

func (r roleRepository) FindByName(name string) (role entity.Role, err error) {
	log.Print("[RoleRepository]...FindByName")
	err = r.DB.Where("name = ?", name).First(&role).Error
	return role, err
}
