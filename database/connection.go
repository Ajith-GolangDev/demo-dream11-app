package database

import (
	"dream_11/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "dream11"
)

var DB *gorm.DB

func ConnectDatabase() {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error

	DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Contest{}, &models.Player{}, &models.UserTeam{})
	if err != nil {
		panic("Failed to create the tables!")
	}
}
