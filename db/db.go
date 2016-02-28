package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := gorm.Open("sqlite3", "lolstat.db")
	if err != nil {
		log.Fatal("Unable to initialize database.")
	}

	db.DB()
}
