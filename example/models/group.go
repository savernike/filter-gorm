package models

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name         string
	PermissionID uint
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
