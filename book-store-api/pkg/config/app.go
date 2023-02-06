package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=john.db.elephantsql.com user=onzcdmyh password=Fz80RSWhh9iwZAU_CZmsdNOBE1_exq2d dbname=onzcdmyh port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}
