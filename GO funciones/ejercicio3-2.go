package main

import "fmt"

func promedio(notas ...float64) float64 {
	var total float64 = 0
	for _, nota := range notas {
		total += nota
	}
	return total / float64(len(notas))
}

func main() {

	fmt.Println("el promedio es: ", promedio(3, 7, 4, 9, 10, 2, 4, 5, 6, 7, 4, 6, 6, 7))
}
