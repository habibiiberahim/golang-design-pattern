package observer

func RunObserver() {
	stockItem := newItem("BELI")

	observerFirst := &Customer{id: "timothy@gmail.com"}
	observerSecond := &Customer{id: "theo@gmail.com"}

	stockItem.register(observerFirst)
	stockItem.register(observerSecond)

	stockItem.updateAvailability()
}
