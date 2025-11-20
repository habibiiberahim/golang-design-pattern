package singleton

import "fmt"

func RunSingleton() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
