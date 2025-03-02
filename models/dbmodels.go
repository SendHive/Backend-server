package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DBSMTPDetails struct {
	Id        uuid.UUID `gorm:"primaryKey,column:id"`
	UserId    uuid.UUID `gorm:"column:user_id;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not_null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not_null"`
	Server    string    `gorm:"column:server;type:varchar(100);not null"`
	Port      string    `gorm:"column:port;type:varchar(100);not null"`
	Username  string    `gorm:"column:username;type:varchar(100);not null"`
	Password  string    `gorm:"column:password;type:varchar(100);not null"`
}

func (DBSMTPDetails) TableName() string {
	return "smtp_tbl"
}

func (*DBSMTPDetails) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}

type DBJobDetails struct {
	Id        uuid.UUID `gorm:"primaryKey,column:id"`
	UserId    uuid.UUID `gorm:"column:user_id;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not_null"`
	TaskId    uuid.UUID `gorm:"column:task_id;not null"`
}

func (DBJobDetails) TableName() string {
	return "job_tbl"
}

func (*DBJobDetails) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New().String()
	tx.Statement.SetColumn("Id", uuid)
	return nil
}
