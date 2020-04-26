package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"api/app/model"
)

func New(uri string) *gorm.DB {
	db, err := gorm.Open("mysql", uri)
	if err != nil {
		fmt.Println("db err: ", err)
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Company{},
		&model.Owner{},
		)
}