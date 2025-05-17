package common

import (
	"github.com/google/uuid"
	"strings"
)

func GetUuids(n int) string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, uuid.New().String())
	}
	return strings.Join(result, "-")
}

func GetRandomStr(len int) string {
	str := GetUuids((len + 36) / 36)
	str = strings.ReplaceAll(str, "-", "")
	return str[:len]
}

func TrimOrPadString(input string, targetLength int) string {
	inputLength := len(input)
	if inputLength > targetLength {
		return input[:targetLength]
	} else if inputLength < targetLength {
		paddingLength := targetLength - inputLength
		padding := strings.Repeat("0", paddingLength)
		return input + padding
	}
	return input
}
