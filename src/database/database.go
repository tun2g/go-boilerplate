package database

import (
	"fist-app/src/apis/model"
	"fmt"
	"os"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	fmt.Println(dataSourceName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
    {
        ID: "20240526000001",
        Migrate: func(tx *gorm.DB) error {
            return tx.AutoMigrate(&model.User{})
        },
        Rollback: func(tx *gorm.DB) error {
            return tx.Migrator().DropTable(&model.User{})
        },
    },
	})
	
	if err := m.Migrate(); err != nil {
    panic("failed to migrate database")
	}

	return db
}
