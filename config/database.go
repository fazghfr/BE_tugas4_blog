package config

import (
	"fmt"

	"tugas4/entitiy"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=passwordOjan dbname=goblog port=5432 TimeZone=Asia/Jakarta"
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
