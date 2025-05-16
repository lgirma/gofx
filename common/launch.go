package common

import (
	"fmt"
	"os/exec"
	"runtime"
)

func LaunchUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin": // macOS
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	command := exec.Command(cmd, args...)
	err := command.Start()
	if err != nil {
		return fmt.Errorf("failed to open browser: %w", err)
	}

	return nil
}
