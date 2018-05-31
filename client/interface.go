package client

type Client interface {
	// Store accepts an id and a payload in bytes and requests that the
	// encryption-server stores them in its data store
	Store(id, payload []byte) (aesKey []byte, err error)

	// Retrieve accepts an id and an AES key, and requests that the
	// encryption-server retrieves the original (decrypted) bytes stored
	// with the provided id
	Retrieve(id, aesKey []byte) (payload []byte, err error)
}
func Store(id, payload []byte) (aesKey []byte, err error) (aesKey []byte, err error) {
	c, err := aes.NewCipher(id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}