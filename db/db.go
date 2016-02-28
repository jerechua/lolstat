package db

import (
	"fmt"
	"log"

	v "github.com/jerechua/validate"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./lolstat.db"

var (
	DB     gorm.DB
	models []interface{}
)

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

	DB = db
}

// Create creates a new model for the given interface.
func Create(i interface{}) error {
	if err := v.Validate(i); err != nil {
		return err
	}
	if err := DB.Create(i).Error; err != nil {
		return fmt.Errorf("Err: %v", err)
	}
	return nil
}
