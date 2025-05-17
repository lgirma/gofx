package licensing

import (
	"fmt"
	"github.com/lgirma/gofx/common"
)

type LicenseIdOptions struct {
	MaxIdLength int
}

const DefaultMaxIdLength = 20

func GetLicenseId(product Product, options LicenseIdOptions) (string, error) {
	if options.MaxIdLength == 0 {
		options.MaxIdLength = DefaultMaxIdLength
	}
	hardwareId, err := GetHardwareId()
	if err != nil {
		return "", err
	}
	productStr := common.TrimOrPadString(product.Name+product.Edition, options.MaxIdLength/2)
	result := fmt.Sprintf("%s%s",
		productStr,
		common.TrimOrPadString(hardwareId, options.MaxIdLength/2))
	return common.TrimOrPadString(result, options.MaxIdLength), nil
}
