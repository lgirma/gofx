//go:build windows

package licensing

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

// GetHardwareId returns the motherboard serial number on Windows.
// It first attempts to use wmic, and if that fails, it uses PowerShell's Get-CimInstance.
func GetHardwareId() (string, error) {
	serial, err := getSerialWmic()
	if err == nil && serial != "" {
		return serial, nil
	}

	serial, err = getSerialPowershell()
	if err == nil && serial != "" {
		return serial, nil
	}

	if err != nil {
		return "", err
	}
	return "", errors.New(ErrHardwareIdEmpty)
}

func getSerialWmic() (string, error) {
	cmd := exec.Command("wmic", "baseboard", "get", "serialnumber")
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.ExitCode() != 0 || bytes.Contains(exitErr.Stderr, []byte("not recognized")) {
				return "", err
			}
		}
		return "", err
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return "", errors.New(ErrHardwareIdEmpty)
	}

	serialNumber := strings.TrimSpace(lines[1])
	return serialNumber, nil
}

// getSerialPowershell attempts to retrieve the serial number using PowerShell's Get-CimInstance.
func getSerialPowershell() (string, error) {
	cmd := exec.Command("powershell.exe", "-Command", "Get-CimInstance win32_baseboard | Select-Object SerialNumber | Format-List")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	lines := strings.SplitSeq(string(output), "\n")
	for line := range lines {
		if strings.Contains(line, "SerialNumber :") {
			parts := strings.SplitN(line, ":", 2) // Split only on the first colon
			if len(parts) > 1 {
				serialNumber := strings.TrimSpace(parts[1])
				return serialNumber, nil
			}
		}
	}

	return "", errors.New(ErrHardwareIdEmpty)
}
