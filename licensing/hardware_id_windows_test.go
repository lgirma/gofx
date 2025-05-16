//go:build windows

package licensing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMotherboardSerial(t *testing.T) {
	serial, err := GetMotherboardSerial()
	assert.NoError(t, err)
	assert.NotEmpty(t, serial)
	t.Logf("Serial: %s", serial)
}
