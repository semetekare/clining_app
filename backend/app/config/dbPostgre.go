package config

import (
	"clining_app/app/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbPostgre *gorm.DB


func ConnectDatabasePostgre() {
	host := os.Getenv("DB_HOST_P")
	user := os.Getenv("DB_USER_P")
	password := os.Getenv("DB_PASSWORD_P")
	dbname := os.Getenv("DB_NAME_P")
	port := os.Getenv("DB_PORT_P")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.User{}, )

	DbPostgre = db
}
