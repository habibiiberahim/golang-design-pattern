package abstract_factory

type IPhone interface {
	setBrand(brand string)
	setStorage(storage int)
	getBrand() string
	getStorage() int
}

type Phone struct {
	brand   string
	storage int
}

func (p *Phone) setBrand(brand string) {
	p.brand = brand
}

func (p *Phone) setStorage(storage int) {
	p.storage = storage
}

func (p *Phone) getBrand() string {
	return p.brand
}

func (p *Phone) getStorage() int {
	return p.storage
}
