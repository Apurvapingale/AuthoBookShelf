package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDB()  {
	d , err := gorm.Open(postgres.Open("user=postgres password=551133 dbname=Books sslmode=disable"))

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
