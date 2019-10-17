package lib

import (
	"os"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Loading mysql dialects
)

var (
	db     *Database
	oncedb sync.Once
)

// Database struct - structure of database
type Database struct {
	*gorm.DB
}

// GetDb - get instance of database
func GetDb() *Database {

	oncedb.Do(func() {

		db = initDb()
	})

	return db
}

func initDb() *Database {
	user := os.Getenv("CONNECTION_USER")
	password := os.Getenv("CONNECTION_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	db, err := gorm.Open("mysql", user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}
