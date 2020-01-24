package ex

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileReader_ReadInvalidData(t *testing.T) {
	r := Reader{}
	config, err := r.Read(strings.NewReader("hello"))
	assert.Error(t, err)
	assert.Nil(t, config)
}
func TestFileReader_ReadValidData(t *testing.T) {
	r := Reader{}
	data := `
encrypter:
  port: 8990
`
	config, err := r.Read(strings.NewReader(data))
	assert.NoError(t, err)
	assert.Equal(t, config.Encrypter.Port, 8990)
}
