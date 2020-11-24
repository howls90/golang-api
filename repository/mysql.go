package repository

import (
	"gorm.io/gorm"

	"os"

	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDB() (err error) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_DATABASE")

	dsn := username + ":" + password + "@tcp(" + host + ":3306)/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err == nil {
		err := DB.AutoMigrate(&User{})
		DB.AutoMigrate(&Post{})
		DB.AutoMigrate(&Comment{})
		return err
	}

	return err
}
