package database

import (
	"fist-app/src/apis/model"
	"fist-app/src/lib/logger"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

var _logger = logger.Logger()

func GetMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "20240526000001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.User{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&model.User{})
			},
		},
	}
}


func Migration() cli.Command{
	cli := cli.Command{
		Name:  "migrate",
		Usage: "Apply all migrations",
		Action: func(c *cli.Context) error {
			db := InitDB()
			m := gormigrate.New(db, gormigrate.DefaultOptions, GetMigrations())
			if err := m.Migrate(); err != nil {
				_logger.Errorf("could not migrate: %s", err)
				return nil
			}
			_logger.Info("Migrations applied successfully!")
			return nil
		},
	}
	return cli
}

func Rollback() cli.Command{
	cli := cli.Command{
		Name:  "rollback",
		Usage: "Rollback the last migration",
		Action: func(c *cli.Context) error {
			db := InitDB()
			m := gormigrate.New(db, gormigrate.DefaultOptions, GetMigrations())
			if err := m.RollbackLast(); err != nil {
				_logger.Errorf("could not rollback: %s", err)
				return nil
			}
			_logger.Info("Migration rolled back successfully!")
			return nil
		},
	}
	return cli;
}