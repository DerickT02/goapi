package models

import (
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func DBConn(){
	database, err := gorm.Open("sqlite3", "test.db")

	if err != nil{
		panic("Not Connected to Database")
	}

	database.AutoMigrate(&Person{})

	DB = database
}