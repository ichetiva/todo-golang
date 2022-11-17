package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserRefer uint
	User      User `gorm:"foreignKey:UserRefer"`
	Token     string
}
