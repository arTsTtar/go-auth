package repository

import (
	"go-auth/entity"
	"gorm.io/gorm"
	"log"
)

type backupCodeRepository struct {
	DB *gorm.DB
}

type BackupCodeRepository interface {
	SaveAll(codes entity.BackupCodes) (entity.BackupCodes, error)
	DeleteById(id uint) error
	FindByUser(user entity.User) (entity.BackupCodes, error)
}

func NewBackupCodeRepository(db *gorm.DB) BackupCodeRepository {
	return backupCodeRepository{
		DB: db,
	}
}

func (b backupCodeRepository) FindByUser(user entity.User) (backupCodes entity.BackupCodes, err error) {
	log.Printf("[BackupCodeRepository]... FindByUser")
	err = b.DB.Where("user_id = ?", user.Id).Find(&backupCodes).Error
	return backupCodes, err
}

func (b backupCodeRepository) SaveAll(backupCodes entity.BackupCodes) (entity.BackupCodes, error) {
	log.Print("[BackupCodeRepository]...SaveAll")
	err := b.DB.Create(&backupCodes).Error
	return backupCodes, err
}

func (b backupCodeRepository) DeleteById(id uint) error {
	log.Print("[BackupCodeRepository]...DeleteById")
	err := b.DB.Delete(&entity.BackupCode{}, id).Error
	return err
}
