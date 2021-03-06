package config

import (
	"database/sql"
	"log"
	"notaku/model"
)

// Datastore test
type Datastore interface {
	AllPosts() ([]*model.Post, error)
}

// DB test
type DB struct {
	*sql.DB
}

// NewDB test
func NewDB() (*DB, error) {
	//connect to db
	var err error
	db, err := sql.Open("sqlite3", "./note.db")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &DB{db}, nil

}
