//go:build darwin

package licensing

import (
	"errors"
	"os/exec"
	"strings"
)

// GetHardwareId returns the motherboard serial number on macOS.
func GetHardwareId() (string, error) {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

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

	return "", errors.New(ErrHardwareIdEmpty)
}
