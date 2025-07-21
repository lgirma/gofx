package common

import (
	"github.com/ericlagergren/decimal"
	"github.com/google/uuid"
	"regexp"
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

func GetRegexGroup(re *regexp.Regexp, str string) map[string]string {
	matches := re.FindStringSubmatch(str)
	groupNames := re.SubexpNames()
	result := make(map[string]string)
	if len(matches) > 0 {
		for i, group := range groupNames {
			if group == "" {
				continue
			}
			result[group] = matches[i]
		}
	}
	return result
}

func GetRegexGroups(re *regexp.Regexp, str string) []map[string]string {
	matches := re.FindAllStringSubmatch(str, -1)
	groupNames := re.SubexpNames()
	var result []map[string]string
	if len(matches) > 0 {
		for _, match := range matches {
			item := map[string]string{}
			for i, group := range groupNames {
				if group == "" {
					continue
				}
				item[group] = match[i]
			}
			result = append(result, item)
		}
	}
	return result
}

func ParseDecimal(str string) *decimal.Big {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, ",", "")
	d := new(decimal.Big)
	d, _ = d.SetString(str)
	return d
}

func IsNullOrEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
