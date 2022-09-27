/*Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Si es categoría C, su salario es de $1.000 por hora
Si es categoría B, su salario es de $1.500 por hora más un %20 de su salario mensual
Si es de categoría A, su salario es de $3.000 por hora más un %50 de su salario mensual

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario.
*/

package main

import "fmt"

func salaryCalculator(minutes float64, category string) float64 {
	var result float64

	if category == "C" {
		result = minutes / 60 * 1000
	}
	if category == "B" {
		result = minutes / 60 * 1500
		result = result + result*0.5
	}
	if category == "A" {
		result = minutes / 60 * 3000
		result = result + result*0.5
	}
	return result
}
func main() {

	//salary := salaryCalculator(2400, "C")
	//salary := salaryCalculator(2400, "B")
	salary := salaryCalculator(2400, "A")

	fmt.Println(salary)
}
