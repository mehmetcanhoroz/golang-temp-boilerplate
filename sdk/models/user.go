package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

const UserModelTableName = "users"

type User struct {
	ID            uint64 `json:"id" gorm:"primaryKey"`
	FullName      string `gorm:"not null"`
	Phone         string `gorm:"unique;not null"`
	Email         string `gorm:"unique;not null"`
	Password      string `gorm:"size:75;not null"`
	Username      string `gorm:"index;unique;not null"`
	Birthday      sql.NullTime
	ActivatedAt   sql.NullTime
	DeactivatedAt sql.NullTime

	CreatedAt time.Time      `gorm:"index"`
	UpdatedAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type GetUserSuccessfulResponse struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Phone    string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Username string `gorm:"index;unique;not null"`
	Birthday sql.NullTime
}

func (User) TableName() string {
	return UserModelTableName
}
