package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Content string
	Done    bool
}
