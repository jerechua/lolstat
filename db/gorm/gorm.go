package gorm

import (
	"fmt"
	"log"

	v "github.com/jerechua/validate"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const dbName = "./lolstat.db"

var (
	models []interface{}
)

func RegisterModel(model interface{}) {
	models = append(models, model)
}

type Gorm struct {
	db gorm.DB
}

func New() *Gorm {
	db, err := gorm.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("Unable to initialize database.")
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// TODO: This is dangerous if there are multiple binaries running live.
	db.AutoMigrate(models...)

	return &Gorm{db}
}

func (g *Gorm) DB() gorm.DB {
	return g.db
}

// Create creates a new model for the given interface.
func (g *Gorm) Create(i interface{}) error {
	if err := v.Validate(i); err != nil {
		return err
	}
	if err := g.db.Create(i).Error; err != nil {
		return fmt.Errorf("Err: %v", err)
	}
	return nil
}

// Exists checks the existence of a record and updates i to be the model in the
// database.
func (g *Gorm) Exists(in interface{}) (bool, error) {
	if err := v.Validate(in); err != nil {
		return false, err
	}
	// TODO: Consider updating/returning i if it exists since we're already
	// querying anyway.
	result := g.db.Where(in).First(in)
	return !result.RecordNotFound(), nil
}
