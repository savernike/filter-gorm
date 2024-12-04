package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name   string
	Groups []Group `gorm:"many2many:user_groups;"` // Tabella intermedia: user_groups
}
