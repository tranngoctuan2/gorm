package database

import (
	"fmt"
	"gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "password"
const DB_NAME = "testdb"
const DB_HOST = "localhost"
const DB_PORT = "3306"

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	DB.AutoMigrate(&models.User{})
	return DB
}

func InitDB() *gorm.DB {
	DB = ConnectDB()
	return DB
}
