package abstract_factory

type Asus struct {
}

func (a *Asus) makeLaptop() ILaptop {
	return &AsusLaptop{
		Laptop{
			brand:   "ASUS",
			storage: 512,
		},
	}
}

func (a *Asus) makePhone() IPhone {
	return &AsusPhone{
		Phone{
			brand:   "ASUS",
			storage: 256,
		},
	}
}
