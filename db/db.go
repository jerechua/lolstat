package db

import (
	"log"

	"./gorm"
	"./mgo"
)

var (
	// The main SQL db
	GORM *gorm.Gorm

	// The match list database in MongoDB.
	MATCH *mgo.Client
)

func Init() {
	GORM = gorm.New()
	log.Print("Connected to GORM")

	MATCH = mgo.New("matchlist", "match")
	log.Print("Connected to MGO")
}
