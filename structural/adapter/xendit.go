package adapter

import "fmt"

type Xendit struct {
}

func (x *Xendit) Charge(amount int) {
	fmt.Println("Xendit: charged", amount)
}
