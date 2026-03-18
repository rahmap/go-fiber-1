package database

import (
	"fiber-rest-api/internal/modules/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&user.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
