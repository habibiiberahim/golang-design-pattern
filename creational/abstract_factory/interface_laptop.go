package abstract_factory

type ILaptop interface {
	setBrand(brand string)
	setStorage(storage int)
	getBrand() string
	getStorage() int
}

type Laptop struct {
	brand   string
	storage int
}

func (l *Laptop) setBrand(brand string) {
	l.brand = brand
}

func (l *Laptop) setStorage(storage int) {
	l.storage = storage
}

func (l *Laptop) getBrand() string {
	return l.brand
}

func (l *Laptop) getStorage() int {
	return l.storage
}
