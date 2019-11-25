package eather

import (
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	db, err := openDb(os.Getenv("DATABASE"))

	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db}
}

func openDb(dialect string) (*gorm.DB, error) {
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	switch dialect {
	case "mysql":
		return gorm.Open(dialect, user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	case "postgres":
		return gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password)
	case "sqlite":
		return gorm.Open("sqlite3", dbname)
	case "mssql":
		return gorm.Open("mssql", "sqlserver://"+user+":"+password+"@"+host+":"+port+"?database="+dbname)
	}

	return nil, errors.New("Dialect not found")
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
