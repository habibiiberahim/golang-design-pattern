package factory

type Swift struct {
	Car
}

func newSwift() ICar {
	return &Swift{
		Car{
			name: "Suzuki Swift",
			year: 2025,
		},
	}
}
