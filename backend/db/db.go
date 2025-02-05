package db

import (
	"fmt"
	"go-get-stuff-done/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

type RawQuery struct {
	Query string
}

var DB Dbinstance

func ConnectDB() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PW"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		log.Fatal("oops no database try again with a database", err)
		os.Exit(2)
	}

	log.Println("yay database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("do the migrations right meow")
	db.AutoMigrate(&models.TodoTask{})

	DB = Dbinstance{
		Db: db,
	}
}
