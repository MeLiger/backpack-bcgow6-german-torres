package main

import "fmt"

func calculadorImpuestos(sueldo float64) float64 {
	var result float64
	if sueldo > 50000 {

		result = sueldo * 0.17

	}
	if sueldo > 150000 {

		result = sueldo*0.10 + sueldo*0.17
	}
	return result
}

func main() {

	impuesto := calculadorImpuestos(160000)
	fmt.Println(impuesto)
}
