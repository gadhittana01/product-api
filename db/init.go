package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gadhittana01/product-api/config"
	"github.com/gadhittana01/product-api/helper"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func InitDB() *gorm.DB {
	config := &config.GlobalConfig{}
	helper.LoadConfig(config)

	dbConn := config.DB

	connString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		dbConn.Host, dbConn.Port, dbConn.User, dbConn.Name,
	)

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB coneected Successfully!")

	return db
}
