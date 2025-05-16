//go:build darwin

package licensing

import (
	"errors"
	"os/exec"
	"strings"
)

// GetMotherboardSerial returns the motherboard serial number on macOS.
// It uses system_profiler and does not require elevated permissions.
func GetMotherboardSerial() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// The output of system_profiler SPHardwareDataType contains a line like:
	// Serial Number (system): [serial_number]
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial Number (system):") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				serialNumber := strings.TrimSpace(parts[1])
				return serialNumber, nil
			}
		}
	}

	return "", errors.New("could not determine serial number")
}
