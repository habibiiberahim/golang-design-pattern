package abstract_factory

import "fmt"

type IGamingFactory interface {
	makeLaptop() ILaptop
	makePhone() IPhone
}

func GetGamingFactory(brand string) (IGamingFactory, error) {
	if brand == "ASUS" {
		return &Asus{}, nil
	}

	if brand == "LENOVO" {
		return &Lenovo{}, nil
	}

	return nil, fmt.Errorf("brand is not available")
}
