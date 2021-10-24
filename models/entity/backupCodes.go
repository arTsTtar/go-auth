package entity

type BackupCode struct {
	Id         uint
	UserId     uint   `gorm:"notnull"`
	BackupCode []byte `json:"-" gorm:"notnull"`
}
