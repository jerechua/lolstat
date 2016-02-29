package db

import (
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
	MATCH = mgo.New("matchlist", "match")
}
