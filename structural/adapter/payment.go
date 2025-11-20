package adapter

type Payment interface {
	Pay(amount float64)
}
