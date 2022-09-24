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

	return nil
}

func Migrate() error {
	var err error

	err = Database.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	err = Database.AutoMigrate(&model.Club{})
	if err != nil {
		return err
	}
	err = Database.AutoMigrate(&model.Aircraft{})
	if err != nil {
		return err
	}
	err = Database.AutoMigrate(&model.Reservation{})
	if err != nil {
		return err
	}
	err = Database.AutoMigrate(&model.Membership{})
	if err != nil {
		return err
	}
	err = Database.AutoMigrate(&model.Ownership{})
	if err != nil {
		return err
	}

	return nil
}

func Cleanup() error {
	var err error

	err = Database.Migrator().DropTable(&model.User{})
	if err != nil {
		return err
	}
	err = Database.Migrator().DropTable(&model.Club{})
	if err != nil {
		return err
	}
	err = Database.Migrator().DropTable(&model.Aircraft{})
	if err != nil {
		return err
	}
	err = Database.Migrator().DropTable(&model.Reservation{})
	if err != nil {
		return err
	}
	err = Database.Migrator().DropTable(&model.Membership{})
	if err != nil {
		return err
	}
	err = Database.Migrator().DropTable(&model.Ownership{})
	if err != nil {
		return err
	}

	return nil
}
