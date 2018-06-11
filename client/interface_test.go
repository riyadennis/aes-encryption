package client

import (
	"testing"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
	"github.com/aes-encryption/middleware"
	"github.com/aes-encryption/models"
	"os"
)

const (
	tableName = "encrypted_data"
	source    = "encrypted_db"
)

func TestRandSeq(t *testing.T) {
	size := 16
	key := randSeq(size)
	assert.Equal(t, 16, len(key))

	size1 := 12
	key1 := randSeq(size1)
	assert.Equal(t, 12, len(key1))
}
func TestEncrypt(t *testing.T) {
	key := randSeq(16)
	textToEncrypt := "Plain text"
	encryptedText, err := encrypt(textToEncrypt, key)
	assert.NoError(t, err)
	assert.NotEmpty(t, encryptedText)
}

func TestDecrypt(t *testing.T) {
	key := randSeq(16)
	textToEncrypt := "Plain text"
	encryptedText, err := encrypt(textToEncrypt, key)
	assert.NoError(t, err)
	decryptedText, err := decrypt(encryptedText, []byte(key))
	assert.NoError(t, err)
	assert.Equal(t, string(decryptedText), textToEncrypt)
}
func TestStore(t *testing.T) {
	id := []byte(randSeq(16))
	payLoad := []byte("plain text")
	config := CreateConfig()
	setupDB(config)
	defer tearDown()
	ac := AesClient{Config: config}
	key, err := ac.Store(id, payLoad, )
	defer tearDown()
	assert.NoError(t, err)
	assert.NotEmpty(t, key)
}
func TestRetrieve(t *testing.T) {
	id := []byte("ssXVlBzgbaiCMRAjWw$$")
	payLoad := []byte("plain text")
	setupDB(CreateConfig())
	config := CreateConfig()
	defer tearDown()
	ac := AesClient{Config: config}
	key, err := ac.Store(id, payLoad)
	defer tearDown()
	assert.NoError(t, err)
	decrypted, err := ac.Retrieve(id, key)
	assert.NoError(t, err)
	assert.Equal(t, string(decrypted), "plain text")
}
func CreateConfig() *middleware.Config {
	db := middleware.Db{Type: "sqlite3", Source: source, User: "root", Password: ""}
	encrypter := middleware.Encryptor{Port: 8083, Db: db}
	return &middleware.Config{
		Encrypter: encrypter,
	}
}
func tearDown() {
	os.Remove(source)
}
func setupDB(config *middleware.Config) {
	CreateTable(config.Encrypter.Db)
}
func CreateTable(db middleware.Db) {
	dbConnec, _ := models.InitDB(db)
	statement, _ := dbConnec.Prepare("CREATE TABLE IF NOT EXISTS " + tableName + "(id varchar(100) NOT NULL PRIMARY KEY,encrypted_text  BLOB,encryption_key varchar(100), InsertedDatetime DATETIME);")
	statement.Exec()
}
