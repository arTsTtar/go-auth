package repository

import (
	"go-auth/models/entity"
	"gorm.io/gorm"
	"log"
)

type backupCodeRepository struct {
	DB *gorm.DB
}

type BackupCodeRepository interface {
	SaveAll(entity.Users) (entity.Users, error)
	Delete(id string) error
}

func NewBackupCodeRepository(db *gorm.DB) BackupCodeRepository {
	return backupCodeRepository{
		DB: db,
	}
}

func (b backupCodeRepository) SaveAll(users entity.Users) (entity.Users, error) {
	log.Print("[BackupCodeRepository]...SaveAll")
	err := b.DB.Create(&users).Error
	return users, err
}

func (b backupCodeRepository) Delete(id string) error {
	log.Print("[BackupCodeRepository]...Delete")
	err := b.DB.Delete("id = ?", id).Error
	return err
}
