package encryption

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptDecryptString(t *testing.T) {
	testCases := []struct {
		name     string
		content  string
		password string
		expect   string // Expected decrypted content
		wantErr  bool   // Whether an error is expected during decryption
	}{
		{
			name:     "Successful Encryption and Decryption",
			content:  "This is a secret message!",
			password: "mySuperSecretPassword123",
			expect:   "This is a secret message!",
			wantErr:  false,
		},
		{
			name:     "Empty Content",
			content:  "",
			password: "password",
			expect:   "",
			wantErr:  false,
		},
		{
			name:     "Empty Password",
			content:  "some content",
			password: "",
			expect:   "some content",
			wantErr:  false,
		},
		{
			name:     "Empty Content and Password",
			content:  "",
			password: "",
			expect:   "",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encrypted, err := EncryptString(tc.content, tc.password)
			assert.NoError(t, err, "EncryptString failed unexpectedly")

			decrypted, err := DecryptString(encrypted, tc.password)

			if tc.wantErr {
				assert.Error(t, err, "DecryptString did not return an error when one was expected")
			} else {
				assert.NoError(t, err, "DecryptString returned an unexpected error")
				assert.Equal(t, tc.expect, decrypted, "Decrypted content mismatch")
			}
		})
	}
}

func TestDecryptString_WrongPassword(t *testing.T) {
	content := "This is a secret message!"
	password := "correctPassword"
	wrongPassword := "incorrectPassword"

	encrypted, err := EncryptString(content, password)
	assert.NoError(t, err)

	_, err = DecryptString(encrypted, wrongPassword)
	assert.Error(t, err)

	if err != nil {
		assert.Equal(t, err.Error(), ErrEncryptionWrongPassword)
	}
}

func TestDecryptString_TamperedData(t *testing.T) {
	content := "Original message"
	password := "password"

	encrypted, err := EncryptString(content, password)
	assert.NoError(t, err)

	// Tamper with the encrypted data (e.g., flip a bit)
	tamperedEncrypted, err := base64.StdEncoding.DecodeString(encrypted)
	assert.NoError(t, err)

	if len(tamperedEncrypted) > 0 {
		tamperedEncrypted[0] = tamperedEncrypted[0] ^ 0xFF // Flip bits of the first byte
	} else {
		t.Skip("Encrypted data is too short to tamper")
	}

	tamperedEncryptedString := base64.StdEncoding.EncodeToString(tamperedEncrypted)

	_, err = DecryptString(tamperedEncryptedString, password)
	assert.Error(t, err)

	// Check if the error message indicates decryption failure
	if err != nil {
		assert.Equal(t, err.Error(), ErrEncryptionWrongPassword)
	}
}

func TestDecryptString_InvalidBase64(t *testing.T) {
	invalidBase64 := "This is not valid base64!"
	password := "password"

	_, err := DecryptString(invalidBase64, password)
	assert.Error(t, err)

	if err != nil {
		assert.Equal(t, err.Error(), ErrEncryptionBase64DecodeFailed)
	}
}

func TestDecryptString_ShortCiphertext(t *testing.T) {
	shortData := make([]byte, 10)
	shortCiphertext := base64.StdEncoding.EncodeToString(shortData)
	password := "password"

	_, err := DecryptString(shortCiphertext, password)
	assert.Error(t, err)

	if err != nil {
		assert.Equal(t, err.Error(), ErrEncryptionCypherTooShort)
	}
}
