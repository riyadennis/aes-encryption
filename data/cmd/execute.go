package cmd

import (
	"github.com/riyadennis/aes-encryption/data/models"
	"github.com/riyadennis/aes-encryption/ex"
	"github.com/sirupsen/logrus"
)

func ExecuteCommand(migrateChoice string, config *ex.Config) bool {
	db, err := models.InitDB(config.Encrypter.Db)
	if err != nil {
		logrus.Errorf("Unable initial  %s", err.Error())
	}
	if migrateChoice == "up" {
		return MigrateUp(db, config.Encrypter.Db.Source)

	}
	if migrateChoice == "down" {
		return MigrateDown(db, config.Encrypter.Db.Source)
	}
	return false
}
