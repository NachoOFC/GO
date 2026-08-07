package main

import "fmt"

func main() {
	// hola mundo GO
	fmt.Println("hola mundo GO de Nacho OFC")

	// VARIABLES
	var nombre string = "Nacho"
	var edad int = 23
	fmt.Println("Hola, mi nombre es", nombre, "y tengo", edad, "años")

	if edad >= 18 && edad <= 100 {
		fmt.Println("Soy mayor de edad")
	} else if edad < 18 && edad >= 0 {
		fmt.Println("Soy menor de edad")
	} else {
		fmt.Println("No se puede determinar si soy mayor o menor de edad")
	}
}
