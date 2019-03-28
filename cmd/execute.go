package cmd

import (
	"github.com/riyadennis/aes-encryption/middleware"
	"github.com/sirupsen/logrus"
	"github.com/riyadennis/aes-encryption/models"
)

func ExecuteCommand(migrateChoice string, config *middleware.Config) (bool) {
	db, err := models.InitDB(config.Encrypter.Db)
	if err != nil {
		logrus.Errorf("Unable initial  %s", err.Error())
	}
	if migrateChoice == "up"{
		return MigrateUp(db, config.Encrypter.Db.Source)

	}
	if migrateChoice == "down"{
		return MigrateDown(db, config.Encrypter.Db.Source)
	}
	return false
}
