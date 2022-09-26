package db

import (
	"github.com/mehmetcanhoroz/digital-marketplace/sdk/models"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewDatabase() *gorm.DB {
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      "golang:example@tcp(127.0.0.1:3306)/digital-marketplace?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize:        256,                                                                                            // default size for string fields
		DisableDatetimePrecision: true,                                                                                           // disable datetime precision, which not supported before MySQL 5.6
		//DontSupportRenameIndex:    true,                                                                                           // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		//DontSupportRenameColumn:   true,                                                                                           // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // autoconfigure based on currently MySQL version
	}), &gorm.Config{})

	dbc, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = dbc.Ping()
	if err != nil {
		panic(err)
	}

	PrepareMigrations()

	return db
}

func PrepareMigrations() {
	err := db.Debug().AutoMigrate(&models.Category{})
	if err != nil {
		panic(err)
	}
	err = db.Debug().AutoMigrate(&models.Item{})
	if err != nil {
		panic(err)
	}
	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
