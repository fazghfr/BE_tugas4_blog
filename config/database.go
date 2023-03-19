package config

import (
	"fmt"
	"os"

	"tugas4/entitiy"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	if err := db.AutoMigrate(
		entitiy.Blog{},
		entitiy.Comment{},
		entitiy.User{},
	); err != nil {
		fmt.Println(err)
		panic("failed to migrate schema")
	}

	return db
}

func Close(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	dbSQL.Close()
}
