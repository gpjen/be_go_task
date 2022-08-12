package config

import (
	"be_go_task/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConn() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/go_task?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.Task{})

	return db, err
}
