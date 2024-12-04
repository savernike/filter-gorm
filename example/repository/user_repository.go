package repository

import (
	"filter-gorm/example/models"
	"filter-gorm/example/models/filter"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUsers(filter filter.UserFilter) ([]models.User, error)
	DeleteUser(user models.User) error
	UpdateUser(id uint, user models.User) error
}
