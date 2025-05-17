package licensing

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
)

func EncryptString(content string, password string) (string, error) {
	key := sha256.Sum256([]byte(password))

	block, err := aes.NewCipher(key[:])
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

	ciphertext := gcm.Seal(nonce, nonce, []byte(content), nil)

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)

	return encodedCiphertext, nil
}

func DecryptString(encryptedString string, password string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", errors.New(ErrEncryptionBase64DecodeFailed)
	}

	key := sha256.Sum256([]byte(password))

	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New(ErrEncryptionCypherTooShort)
	}

	nonce, sealedCiphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, sealedCiphertext, nil)
	if err != nil {
		return "", errors.New(ErrEncryptionWrongPassword)
	}

	return string(plaintext), nil
}
