package abstract_factory

type Lenovo struct {
}

func (l *Lenovo) makeLaptop() ILaptop {
	return &LenovoLaptop{
		Laptop{
			brand:   "LENOVO",
			storage: 512,
		}}
}

func (l *Lenovo) makePhone() IPhone {
	return &LenovoPhone{
		Phone{
			brand:   "LENOVO",
			storage: 128,
		},
	}
}
