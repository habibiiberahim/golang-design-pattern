package factory

type ICar interface {
	setName(name string)
	setYear(year int)
	getName() string
	getYear() int
}
