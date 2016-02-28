package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./lolstat.db"

var models []interface{}

func RegisterModel(model interface{}) {
	models = append(models, model)
}

func Init() {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("Unable to initialize database.")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(models...)
}
