package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}
