package factory

import "fmt"

func RunFactory() {
	swift, _ := getCar("swift")
	brio, _ := getCar("brio")

	printDetails(swift)
	printDetails(brio)
}

func printDetails(c ICar) {
	fmt.Printf("Car Name : %s", c.getName())
	fmt.Println()
	fmt.Printf("Car Year : %d", c.getYear())
	fmt.Println()
}
