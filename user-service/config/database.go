package config

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // dummy import mysql driver
)

var (
	dbHost          = Get("DB_HOST")
	dbPort          = Get("DB_PORT")
	dbName          = Get("DB_NAME")
	dbUser          = Get("DB_USER")
	dbPassword      = Get("DB_PASSWORD")
	dbMaxConnection = GetInt("DB_MAX_CONNECTION")
	dbMinConnection = GetInt("DB_MIN_CONNECTION")
)

var once sync.Once
var db *gorm.DB

// getMySQLURL -> builds mysql url for mysql
func getMySQLURL() string {
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
}

// instantiateDB -> instantiates DB
func instantiateDB() *gorm.DB {
	dbInstance, err := gorm.Open("mysql", getMySQLURL())

	if err != nil {
		panic("Unable to Connect to DB" + err.Error())
	}

	dbInstance.DB().SetMaxIdleConns(dbMinConnection)
	dbInstance.DB().SetMaxOpenConns(dbMaxConnection)

	db = dbInstance
	return db
}

// StartDB -> Configures db and opens db connection
func StartDB() {
	once.Do(func() {
		instantiateDB()
	})
}

// GetDB -> Configures db and opens db connection
func GetDB() *gorm.DB {
	return db
}
