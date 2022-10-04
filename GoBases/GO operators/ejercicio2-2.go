package main

import "fmt"

func main() {

	switch edad := 27; {
	case edad < 22:
		fmt.Println("No cumples con el requisito de edad")
	default:
		fmt.Println("Cumples con el requisito de edad")

	}
	switch antiguedad := 4; {
	case antiguedad < 2:
		fmt.Println("No posees la antiguedad necesaria")
	default:
		fmt.Println("Cumples con la condición de antiguedad")

	}
	switch sueldo := 170000; {
	case sueldo < 100000:
		fmt.Println("Se realizará cobro de intereses")
	default:
		fmt.Println("No se cobrarán intereses")

	}

}
