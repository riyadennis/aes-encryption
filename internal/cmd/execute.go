package cmd

import (
	"github.com/riyadennis/aes-encryption/internal"
	"github.com/riyadennis/aes-encryption/internal/models"
	"github.com/sirupsen/logrus"
)

func ExecuteCommand(migrateChoice string, config *internal.Config) bool {
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
