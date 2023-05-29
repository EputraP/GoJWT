package models

import (
	"gorm.io/gorm"
)

type UserList struct {
	UserId   uint   `gorm:"primarykey"`
	Name     string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(50);not null"`
	Role     string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50);not null"`
	gorm.Model
}
