package adapter

func RunAdapter() {
	paymentGateway := &Xendit{}
	payment := &XenditAdapter{gateway: paymentGateway}

	payment.Pay(100000)
}
