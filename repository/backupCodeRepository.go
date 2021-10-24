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
	SaveAll(codes entity.BackupCodes) (entity.BackupCodes, error)
	Delete(id string) error
}

func NewBackupCodeRepository(db *gorm.DB) BackupCodeRepository {
	return backupCodeRepository{
		DB: db,
	}
}

func (b backupCodeRepository) SaveAll(backupCodes entity.BackupCodes) (entity.BackupCodes, error) {
	log.Print("[BackupCodeRepository]...SaveAll")
	err := b.DB.Create(&backupCodes).Error
	return backupCodes, err
}

func (b backupCodeRepository) Delete(id string) error {
	log.Print("[BackupCodeRepository]...Delete")
	err := b.DB.Delete("id = ?", id).Error
	return err
}
