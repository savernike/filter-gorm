package repository

import (
	"filter-gorm/example/models"
	"filter-gorm/example/models/filter"
)

type GroupRepository interface {
	CreateGroup(Group *models.Group) error
	GetGroups(filter filter.GroupFilter) ([]models.Group, error)
	DeleteGroup(Group models.Group) error
	UpdateGroup(id uint, Group models.Group) error
}
