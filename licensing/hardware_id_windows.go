//go:build windows

package licensing

import (
	"os/exec"
	"strings"
)

// GetMotherboardSerial returns the motherboard serial number on Windows.
// It attempts to use wmic and does not require elevated permissions.
func GetMotherboardSerial() (string, error) {
	cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// The output of wmic is typically in the format:
	// SerialNumber
	// [serial_number]
	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return "", errors.New("could not determine serial number")
	}

	serialNumber := strings.TrimSpace(lines[1])
	return serialNumber, nil
}
