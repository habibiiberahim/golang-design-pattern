package abstract_factory

import "fmt"

func RunAbstractFactory() {
	asusFactory, _ := GetGamingFactory("ASUS")
	lenovoFactory, _ := GetGamingFactory("LENOVO")

	asusLaptop := asusFactory.makeLaptop()
	asusPhone := asusFactory.makePhone()

	lenovoLaptop := lenovoFactory.makeLaptop()
	lenovoPhone := lenovoFactory.makePhone()

	printLaptopDetails(asusLaptop)
	printPhoneDetails(asusPhone)

	printLaptopDetails(lenovoLaptop)
	printPhoneDetails(lenovoPhone)
}

func printLaptopDetails(l ILaptop) {
	fmt.Printf("Brand Logo: %s", l.getBrand())
	fmt.Println()
	fmt.Printf("Storage Size: %d", l.getStorage())
	fmt.Println()
}

func printPhoneDetails(p IPhone) {
	fmt.Printf("Brand Logo: %s", p.getBrand())
	fmt.Println()
	fmt.Printf("Storage Size: %d", p.getStorage())
	fmt.Println()
}
