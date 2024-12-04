package repository

import (
	"github.com/R3n3r0/filter-gorm/example/models"
	"github.com/R3n3r0/filter-gorm/example/models/filter"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers(filter filter.UserFilter) ([]models.User, error)
	DeleteUser(user models.User) error
	UpdateUser(id uint, user models.User) error
}
