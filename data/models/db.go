package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/riyadennis/aes-encryption/ex"
	"github.com/sirupsen/logrus"
)

func InitDB(db ex.Db) (*sql.DB, error) {
	//for mysql
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(mysql-development:3306)/%s?multiStatements=true", db.User, db.Password, db.Source)
	if db.Type == "sqlite3" {
		dbConnectionString = db.Source
	}
	dbConnector, err := sql.Open(db.Type, dbConnectionString)
	if err != nil {
		logrus.Errorf("unable to start database :: %v", err)
		return nil, err
	}
	return dbConnector, nil
}

func getCurrentTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
