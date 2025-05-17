package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadString(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}
	input = strings.TrimRight(input, "\r\n")
	return input, nil
}
