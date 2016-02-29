package db

import (
	"./gorm"
)

var (
	GORM *gorm.Gorm
)

func Init() {
	GORM = gorm.New()
}
