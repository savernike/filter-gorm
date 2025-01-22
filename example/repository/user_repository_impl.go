package repository

import (
	"github.com/R3n3r0/filter-gorm/example/models"
	"github.com/R3n3r0/filter-gorm/example/models/filter"
	"github.com/R3n3r0/filter-gorm/filter_helper"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db            *gorm.DB
	filterService filter_helper.FilterService
}

func (u UserRepositoryImpl) CreateUser(user *models.User) error {
	return u.db.Create(user).Error

}

func (u UserRepositoryImpl) GetUsers(filter filter.UserFilter) ([]models.User, error) {
	query := u.filterService.CreateFilter(filter, &models.User{})
	var us []models.User
	err := query.Preload("Groups").Find(&us).Error
	if err != nil {
		return nil, err
	}

	return us, nil
}

func (u UserRepositoryImpl) DeleteUser(user models.User) error {
	//TODO implement me
	return u.db.Delete(&user).Error
}

func (u UserRepositoryImpl) UpdateUser(id uint, user models.User) error {
	return u.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func NewUserRepositoryImpl(db *gorm.DB, filterService filter_helper.FilterService) UserRepository {
	return &UserRepositoryImpl{db, filterService}
}
