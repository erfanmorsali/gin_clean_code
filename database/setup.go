package database

import (
	"clean_code/cars/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDataBase() (*gorm.DB, error) {
	dsn := "root:12654778@tcp(127.0.0.1:3306)/gin_practice"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.Car{}); err != nil {
		return nil, err
	}
	return db, nil
}
