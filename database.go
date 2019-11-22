package eather

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
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
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")

	db, err := gorm.Open("mysql", user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}

// ModelBase contains common columns for all tables.
type ModelBase struct {
	DatabaseID
	DatabaseCreatedAt
	DatabaseUpdatedAt
	DatabaseDeletedAt
}

// DatabaseID set default ID column
type DatabaseID struct {
	ID uuid.UUID `gorm:"primary_key" json:",omitempty"`
}

// DatabaseCreatedAt set default created_at column
type DatabaseCreatedAt struct {
	CreatedAt time.Time `json:",omitempty"`
}

// DatabaseUpdatedAt set default updated_at column
type DatabaseUpdatedAt struct {
	UpdatedAt time.Time `json:",omitempty"`
}

// DatabaseDeletedAt set default deleted_at column
type DatabaseDeletedAt struct {
	DeletedAt *time.Time `sql:"index" json:",omitempty"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *DatabaseID) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewRandom()

	if err != nil {
		log.Println("Error creating new UUID")
	}

	return scope.SetColumn("ID", uuid)
}
