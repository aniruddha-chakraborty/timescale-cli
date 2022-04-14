package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Database struct {
	postgres_connection *gorm.DB
	Config *Config
}

func (DB *Database) PostgresConnect() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s binary_parameters=yes",
		DB.Config.Get("TIMESCALE","HOST"),
		DB.Config.Get("TIMESCALE","USER"),
		DB.Config.Get("TIMESCALE","PASSWORD"),
		DB.Config.Get("TIMESCALE","DATABASE"),
		DB.Config.Get("TIMESCALE","PORT")))
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(2)
	db.DB().SetConnMaxLifetime(time.Nanosecond)

	if err != nil {
		panic(err)
	}
	fmt.Println("postgres database connected")
	DB.postgres_connection = db
}

func (DB *Database) PostgresConnection() *gorm.DB {
	return DB.postgres_connection
}