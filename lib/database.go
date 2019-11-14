package lib

import (
	"os"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Loading mysql dialects
	uuid "github.com/satori/go.uuid"
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

// Base contains common columns for all tables.
type Base struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	return scope.SetColumn("ID", uuid)
}
