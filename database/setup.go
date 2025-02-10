package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go_fiber_1?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	println("Database connected successfully!")

	DB = database
}
