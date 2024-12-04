package repository

import (
	"filter-gorm/example/models"
	"filter-gorm/example/models/filter"
	"filter-gorm/filter_helper"
	"gorm.io/gorm"
)

type GroupRepositoryImpl struct {
	db            *gorm.DB
	filterService filter_helper.FilterService
}

func (u GroupRepositoryImpl) CreateGroup(Group *models.Group) error {
	return u.db.Create(Group).Error

}

func (u GroupRepositoryImpl) GetGroups(filter filter.GroupFilter) ([]models.Group, error) {
	query := u.filterService.CreateFilter(filter, models.Group{})
	var us []models.Group
	err := query.Find(&us).Error
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (u GroupRepositoryImpl) DeleteGroup(Group models.Group) error {
	//TODO implement me
	return u.db.Delete(&Group).Error
}

func (u GroupRepositoryImpl) UpdateGroup(id uint, Group models.Group) error {
	return u.db.Model(&models.Group{}).Where("id = ?", id).Updates(Group).Error
}

func NewGroupRepositoryImpl(db *gorm.DB, filterService filter_helper.FilterService) GroupRepository {
	return &GroupRepositoryImpl{db, filterService}
}
