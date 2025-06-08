package db

import (
	"os"

	"github.com/RyuichiroYoshida/quest-board-project/models"
	"github.com/RyuichiroYoshida/quest-board-project/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDb() *gorm.DB {
	// DB接続
	dsn := getDSN()
	if dsn == "" {
		utils.LogError("Failed to get database DSN.")
		return nil
	}
	// データベース接続
	database, err := connectDB(dsn)
	if err != nil {
		utils.LogError("Failed to connect to the database: %v", err)
	}
	// マイグレーション
	database.AutoMigrate(&models.User{}, &models.Posts{}, &models.Applications{})
	return database
}

func connectDB(dsn string) (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	if host == "" || user == "" || password == "" || port == "" || dbname == "" {
		utils.LogError("Database environment variables are not set properly.")
		return ""
	}
	return "host=" + host + " user=" + user + " password=" + password + " port=" + port + " dbname=" + dbname + " sslmode=disable TimeZone=Asia/Tokyo"
}
