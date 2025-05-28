package utils

import (
	"addressbook/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

//IF WE WANT TO HANDLE THE ERROR
// func ConnectDB(config *config.Config) (*gorm.DB, error) {
// 	dsn :=fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
// 	config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName, config.DBSSLMode)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err // 👉 Let the caller (main.go) handle it
// 	}
// 	return db, nil
// }
