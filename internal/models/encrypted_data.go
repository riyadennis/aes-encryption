package models

import (
	"github.com/riyadennis/aes-encryption/internal"
	"github.com/sirupsen/logrus"
)

const tableName = "encrypted_data"

type Data struct {
	EncryptedText string
}

func SavePayload(id, key string, payLoad []byte, confDb internal.Db) error {
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

func GetPayLoad(id string, confDb internal.Db) (*Data, error) {
	var encrypted_text string
	var data Data
	db, err := InitDB(confDb)
	defer db.Close()
	if err != nil {
		logrus.Errorf("Unable to get data from db %s", err.Error())
		return nil, err
	}
	query := "SELECT encrypted_text from " + tableName + " where id = '" + id + "'"
	row := db.QueryRow(query)
	err = row.Scan(&encrypted_text)
	if err != nil {
		logrus.Errorf("Unable to get image details %s", err.Error())
		return nil, err
	}
	data.EncryptedText = encrypted_text
	return &data, nil
}
