package auth

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
    ID       string `gorm:"primaryKey"`
    Username string
    Email    string `gorm:"unique"`
    Password string
}

func InitDB() error {
	var err error

	DB, err = gorm.Open(sqlite.Open("users db"), &gorm.Config{})
	if err != nil {
		return err
	}

	return DB.AutoMigrate(&User{})
}
