package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/michaelgbenle/fintech/internal/ports"
)

type Postgres struct {
	DB *gorm.DB
}

//NewDB create/returns a new instance of our Database
func NewDB(DB *gorm.DB) ports.Repository {
	return &Postgres{
		DB: DB,
	}
}

//Initialize opens the database, create tables if not created and populate it if its empty and returns a DB
func Initialize(dbURI string) (*gorm.DB, error) {

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal( err)

	}
	conn.AutoMigrate()

	return conn, nil
}