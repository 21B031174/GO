package database

import (
	"fmt"
	"github.com/21B031174/GolandProjects/assignment_3/model"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=db port=5432 user=admin password=admin dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&model.Book{})
	return db
}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep = time.Duration(1)
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("Database is unavailable %d sec.\n", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}
	return dbase
}
