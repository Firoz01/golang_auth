package lib
    
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func ensureKeyLength(key string) string {
	if len(key) < 16 {
		return key + string(make([]byte, 16-len(key)))
	}
	return key[:16]
}

func Encrypt(data, passphrase string) (string, error) {
	key := ensureKeyLength(passphrase)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedData, passphrase string) (string, error) {
	data, err := base64.URLEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", err
	}

	key := ensureKeyLength(passphrase)

	fmt.Println(".......key......", key)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
