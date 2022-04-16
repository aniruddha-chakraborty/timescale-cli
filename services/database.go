package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Database struct {
	postgres_connection *gorm.DB
	Config *Config
}

func (DB *Database) PostgresConnect() {

	db, err := gorm.Open("postgres", "host=0.0.0.0 port=5435 user=postgres dbname=homework password=password sslmode=disable")
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(time.Nanosecond)

	if err != nil {
		panic(err)
	}
	DB.postgres_connection = db
}

func (DB *Database) PostgresConnection() *gorm.DB {
	return DB.postgres_connection
}