package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup : initializing mysql database
func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("host")
	port := os.Getenv("port")
	user := os.Getenv("user")
	dbname := os.Getenv("dbname")
	password := os.Getenv("password")
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)), &gorm.Config{})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(0)

	err = sqlDB.Ping()
	if err != nil {
		err = fmt.Errorf("error pinging db: %w", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// type Config struct {
// 	Host string `yaml:"server_host"`
// }

// func InitConfig() Config {
// 	var cfg Config
// 	cfg.Host = os.Getenv("server_host")
// 	return cfg
// }
