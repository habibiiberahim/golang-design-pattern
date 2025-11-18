package factory

type Car struct {
	name string
	year int
}

func (c *Car) setName(name string) {
	c.name = name
}

func (c *Car) setYear(year int) {
	c.year = year
}

func (c *Car) getName() string {
	return c.name
}

func (c *Car) getYear() int {
	return c.year
}
