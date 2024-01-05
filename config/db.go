package config

import (
	"github.com/rohit-jaisinghani/gin-gorm-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect func is used to connect with postgres sql.
func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//will map with artist struct and create the table.
	db.AutoMigrate(&models.Artist{})
	DB = db
}
