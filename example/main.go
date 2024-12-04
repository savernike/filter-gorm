package main

import (
	"filter-gorm/example/models"
	"filter-gorm/example/models/filter"
	"filter-gorm/example/repository"
	"filter-gorm/filter_helper"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func printGroup(group []models.Group) {
	for _, g := range group {
		fmt.Println(g.Name)
	}
}
func printUsers(users []models.User) {
	for _, user := range users {
		fmt.Println(user.Name)
		fmt.Println("GROUP")
		for _, g := range user.Groups {
			fmt.Println(g.ID)
		}
		fmt.Println("END GROUP")
		fmt.Println(user.ID)
		fmt.Println(user.CreatedAt)
	}
}
func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.User{}, &models.Group{}, &models.Permission{})
	if err != nil {
		panic(err)
	}
	filterService := filter_helper.NewFilterService(db)
	userRepository := repository.NewUserRepositoryImpl(db, filterService)
	groupRepository := repository.NewGroupRepositoryImpl(db, filterService)
	permissions := []models.Permission{
		{Name: "test"},
		{Name: "video"},
		{Name: "audio"},
		{Name: "file"},
	}
	groups := []models.Group{
		{Name: "group1", Permission: permissions[0]},
		{Name: "group2", Permission: permissions[1]},
		{Name: "group3", Permission: permissions[2]},
		{Name: "group4", Permission: permissions[3]},
	}
	users := []models.User{
		{Name: "user1", Groups: groups},
		{Name: "user2", Groups: groups[2:3]},
		{Name: "user3", Groups: groups[3:4]},
		{Name: "user14", Groups: groups[1:2]},
	}

	for _, user := range users {
		err := userRepository.CreateUser(&user)
		if err != nil {
			panic(err)
		}
	}

	/*userFilter := filter.UserFilter{
		Name:      "user1",
		SortBy:    "ID",
		SortOrder: "ASC",
		Page:      0,
		Size:      10,
	}
	getUsers, err := userRepository.GetUsers(userFilter)
	if err != nil {
		panic(err)
	}
	printUsers(getUsers)
	userFilter2 := filter.UserFilter{
		Search:    "user1",
		SortBy:    "ID",
		SortOrder: "ASC",
		Page:      0,
		Size:      10,
	}

	getUsers, err = userRepository.GetUsers(userFilter2)
	if err != nil {
		panic(err)
	}
	printUsers(getUsers)*/
	ff := []uint{1, 2}
	userFilter3 := filter.UserFilter{
		Groups:    ff,
		SortBy:    "ID",
		SortOrder: "ASC",
		Page:      0,
		Size:      10,
	}

	fmt.Println("START FILTER FOR GROUP RELATED")
	getUsers, err := userRepository.GetUsers(userFilter3)
	if err != nil {
		panic(err)
	}
	printUsers(getUsers)

	groupFilter3 := filter.GroupFilter{
		Permission: "audio",
	}
	getGroups, err := groupRepository.GetGroups(groupFilter3)
	if err != nil {
		panic(err)
	}
	printGroup(getGroups)
}
