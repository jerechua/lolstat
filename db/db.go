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

// Exists checks the existence of a record and updates i to be the model in the
// database.
func Exists(in interface{}) (bool, error) {
	if err := v.Validate(in); err != nil {
		return false, err
	}
	// TODO: Consider updating/returning i if it exists since we're already
	// querying anyway.
	var i interface{}
	return !DB.Where(in).First(i).RecordNotFound(), nil
}
