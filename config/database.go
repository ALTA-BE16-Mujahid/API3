package config

import (
	"fmt"

	"github.com/mujahxd/api3-jwt/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func ConnectionDB(config *Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)

	db, err := gorm.Open(mysql.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	fmt.Println("Connected Successfully to the database")
	return db
}
