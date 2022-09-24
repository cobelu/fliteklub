package config

import (
	"fliteklub/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DatabaseUri = "flite-klub.db"
var Database *gorm.DB

func Connect() error {
	var err error

	Database, err = gorm.Open(sqlite.Open(DatabaseUri), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	err = Migrate()
	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	var err error

	err = Database.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}

	return nil
}
