package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	return nil
}

func DB() *gorm.DB {
	return db
}
