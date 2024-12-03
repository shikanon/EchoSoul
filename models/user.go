package models

import (
	"gorm.io/gorm"
)

// User 表示APP用户
// gorm.Model 会自动带有 ID、CreatedAt、UpdatedAt 和 DeletedAt 字段。
type User struct {
	gorm.Model
	PhoneNumber string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
