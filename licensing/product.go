package licensing

type Product struct {
	Name    string
	Edition string
}

type LicensedProductProvider interface {
	GetProduct() Product
}

type StaticLicensedProductProvider struct {
	info Product
}

func (s *StaticLicensedProductProvider) GetProduct() Product {
	return s.info
}

func NewStaticLicensedProductProvider(info Product) LicensedProductProvider {
	return &StaticLicensedProductProvider{info: info}
}
