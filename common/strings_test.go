package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimOrPadString(t *testing.T) {
	result := TrimOrPadString("abc", 1)
	assert.Equal(t, "a", result)

	result = TrimOrPadString("abc", 5)
	assert.Equal(t, "abc00", result)
}

func TestGetRandomString(t *testing.T) {
	rand1 := GetRandomStr(5)
	assert.Len(t, rand1, 5)
	rand2 := GetRandomStr(5)
	assert.Len(t, rand2, 5)
	assert.NotEqual(t, rand1, rand2)
	rand3 := GetRandomStr(7)
	assert.Len(t, rand3, 7)
}
