package main

import "fmt"

func main() {
	println("Escreva sua temperatura: ")
	var temperaturaCelsius float64
	fmt.Scanln(&temperaturaCelsius)

	// O operador curto apenas pode usado dentro dos codeblocks e não em nível de classe ou global.
	temperaturaFahrenheit := (9*(temperaturaCelsius/5) + 32)
	temperaturaKelvin := (temperaturaCelsius + 273)
	fmt.Println("A temperatura em Celsius é: ", temperaturaKelvin, "K")
	fmt.Printf("Digite a temperatura em Fahrenheit é %g e  temperatura em Celsius é: %g .", temperaturaFahrenheit, temperaturaCelsius)

}
