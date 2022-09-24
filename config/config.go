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

	db := sqlite.Open(DatabaseUri)
	Database, err = gorm.Open(db, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func Migrate(models []model.Model) error {
	var err error

	for _, m := range models {
		err = Database.AutoMigrate(&m)
		if err != nil {
			return err
		}
	}

	return nil
}

func Cleanup(models []model.Model) error {
	var err error

	for _, m := range models {
		err = Database.Migrator().DropTable(&m)
		if err != nil {
			return err
		}
	}

	return nil
}
