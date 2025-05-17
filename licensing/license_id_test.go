package licensing

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLicenseId(t *testing.T) {
	product := Product{
		Name:    "product",
		Edition: "edition",
	}
	id, err := GetLicenseId(product, LicenseIdOptions{})
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
	assert.Len(t, id, DefaultMaxIdLength)
	t.Logf("Id: %s", id)

	id2, err := GetLicenseId(product, LicenseIdOptions{MaxIdLength: 8})
	assert.Nil(t, err)
	assert.NotEmpty(t, id2)
	assert.Len(t, id2, 8)
}
