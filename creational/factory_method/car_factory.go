package factory_method

import (
	"fmt"
)

func getCar(carType string) (ICar, error) {
	if carType == "swift" {
		return newSwift(), nil
	}
	if carType == "brio" {
		return newBrio(), nil
	}

	return nil, fmt.Errorf("unknown car type")
}
