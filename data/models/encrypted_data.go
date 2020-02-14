package models

import (
	"github.com/riyadennis/aes-encryption/ex"
	"github.com/sirupsen/logrus"
)

const tableName = "encrypted_data"

type Data struct {
	EncryptedText string
}

func SavePayload(id, key string, payLoad []byte, confDb ex.Db) error {
	db, err := InitDB(confDb)
	defer db.Close()
	if err != nil {
		logrus.Errorf("Unable to save payload %s", err.Error())
		return err
	}
	query, err := db.Prepare("INSERT INTO " + tableName + "(id,encrypted_text,encryption_key,InsertedDatetime) VALUES(?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("Unable to save payload %s", err.Error())
		return err
	}
	_, err = query.Exec(id, payLoad, key, getCurrentTimeStamp())
	if err != nil {
		logrus.Errorf("Unable to save payload %s", err.Error())
		return err
	}
	return nil
}

func GetPayLoad(id string, confDb ex.Db) (*Data, error) {
	var encryptedText string
	var data Data
	db, err := InitDB(confDb)
	defer db.Close()
	if err != nil {
		logrus.Errorf("failed to initialise database :: %v", err)
		return nil, err
	}
	query := "SELECT encrypted_text from " + tableName + " where id = '" + id + "'"
	row := db.QueryRow(query)
	err = row.Scan(&encryptedText)
	if err != nil {
		logrus.Errorf("failed to fetch data from table :: %v", err)
		return nil, err
	}
	data.EncryptedText = encryptedText
	return &data, nil
}
