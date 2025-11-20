package adapter

type XenditAdapter struct {
	gateway *Xendit
}

func (x *XenditAdapter) Pay(amount float64) {
	x.gateway.Charge(int(amount))
}
