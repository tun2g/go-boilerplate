package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" //nolint
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	fmt.Println(dataSourceName)

	db, err := gorm.Open(os.Getenv("DATABASE_DRIVER"), dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DATABASE_MAX_OPEN_CONNS"))

	if err != nil {
		log.Fatal(err)
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DATABASE_MAX_IDLE_CONNS"))

	if err != nil {
		log.Fatal(err)
	}

	connMaxLife, err := strconv.Atoi(os.Getenv("DATABASE_CONN_MAX_LIFE"))

	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxOpenConns(maxOpenConns)
	db.DB().SetMaxIdleConns(maxIdleConns)
	db.DB().SetConnMaxLifetime(time.Duration(connMaxLife) * time.Second)

	return db
}
