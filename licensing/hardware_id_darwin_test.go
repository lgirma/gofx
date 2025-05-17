//go:build darwin

package licensing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHardwareId(t *testing.T) {
	serial, err := GetHardwareId()
	assert.NoError(t, err)
	assert.NotEmpty(t, serial)
	t.Logf("Serial: %s", serial)
}

func TestGetHardwareIdNotRandom(t *testing.T) {
	serial, err := GetHardwareId()
	assert.NoError(t, err)
	serial2, err := GetHardwareId()
	assert.NoError(t, err)
	assert.Equal(t, serial, serial2)
}
