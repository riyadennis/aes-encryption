package cmd

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/sirupsen/logrus"
)

const (
	//when we add a new migration this constant need to be updated
	step = 1
	//if we change the folder in which we keep our migration
	//files we need to update this
	sourceUrl = "file://migrations/"
)

func MigrateUp(db *sql.DB, databaseName string) bool {
	fmt.Println("Running migrations ..")
	m := setUpForMigration(db, databaseName)
	err := m.Steps(step)
	if err != nil {
		logrus.Errorf("unable to run migration :: %v", err)
		return false
	}
	fmt.Println("Done")
	return true
}
func setUpForMigration(db *sql.DB, databaseName string) *migrate.Migrate {
	migrationConfig := &mysql.Config{}
	driver, _ := mysql.WithInstance(db, migrationConfig)
	m, err := migrate.NewWithDatabaseInstance(
		sourceUrl,
		databaseName,
		driver,
	)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	return m
}
func MigrateDown(db *sql.DB, databaseName string) bool {
	fmt.Println("Undoing  migrations")
	m := setUpForMigration(db, databaseName)
	err := m.Down()
	if err != nil {
		logrus.Errorf("unable to run migration down :: %v", err)
		return false
	}
	fmt.Println("Done")
	return true
}
