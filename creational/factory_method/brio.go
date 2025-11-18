package factory_method

type Brio struct {
	Car
}

func newBrio() ICar {
	return &Brio{
		Car{
			name: "Honda Brio",
			year: 2025,
		},
	}
}
