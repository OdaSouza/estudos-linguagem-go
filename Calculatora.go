package main

import "fmt"

func main() {
	fmt.Println("Números de 1 a 100 divisíveis por 3:")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println() // Nova linha no final
}
