package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ModifiledAt string `json:"modifiled_at"`
}

//Create a new user
func CreateUser(db *gorm.DB, user *User) (err error) {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

//Get all users
func GetUsers(db *gorm.DB, user *[]User) (err error) {
	if err := db.Find(&user).Error; err != nil {
		return err
	}
	return nil
}

//Update a user
func UpdateUser(db *gorm.DB, user *User, id int32) (err error) {
	if err := db.Where("id =?", id).Save(user).Error; err != nil {
		return err
	}
	return nil
}

//Delete a user
func DeleteUser(db *gorm.DB, user *User, id int32) (err error) {
	if err := db.Where("id = ?", id).Delete(user).Error; err != nil {
		return err
	}
	return nil
}

//Get a user by ID
func GetUserByID(db *gorm.DB, user *User, id int32) (err error) {
	if err := db.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
