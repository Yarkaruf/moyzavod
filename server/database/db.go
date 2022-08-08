package database

import (
	"myzavod/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB - Global gorm handle
var DB *gorm.DB

// Init gorm session
func Init() (err error) {
	// TODO: Use config struct for this
	dsn := "host=localhost user=postgres password=postgres919191 dbname=moyzavod port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Location{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.Service{})
	DB.AutoMigrate(&models.Technology{})
	DB.AutoMigrate(&models.Variation{})
	DB.AutoMigrate(&models.Params{})

	DB.FirstOrCreate(&models.Service{
		Cost: 300,
	})

	return nil
}
