package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			possiblePaths := []string{".env", "../../.env", "../.env"}
			loaded := false

			for _, path := range possiblePaths {
				if err := godotenv.Load(path); err == nil {
					loaded = true
					break
				}
			}

			if !loaded {
				fmt.Println("Warning: Could not find .env file. Using environment variables directly.")
			}
		}
	}

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&tls=true",
		user, pass, host, port, name)

	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		fmt.Println("Error: Missing database configuration. Please check your environment variables.")
		fmt.Println("Required variables: DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE")
		panic("Database configuration missing")
	}

	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
