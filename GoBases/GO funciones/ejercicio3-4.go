/*Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas
de calificaciones de los alumnos de un curso, requiriendo calcular los valores mínimo,
 máximo y promedio de sus calificaciones.

Se solicita generar una función que indique qué tipo de cálculo se quiere
realizar (mínimo, máximo o promedio) y que devuelva otra función
( y un mensaje en caso que el cálculo no esté definido)
que se le puede pasar una cantidad N de enteros y devuelva el cálculo
que se indicó en la función anterior
Ejemplo: */

package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func minFunc(values ...float64) float64 {
	var min float64
	for i, value := range values {
		if i == 0 {
			min = value
		}
		if value < min {
			min = value
		}
	}
	return min
}

func maxFunc(values ...float64) float64 {
	var max float64
	for i, value := range values {
		if i == 0 {
			max = value
		}
		if value > max {
			max = value
		}
	}
	return max
}

func avgFunc(values ...float64) float64 {
	var avg float64
	var totalNotas float64

	for _, value := range values {
		totalNotas += value // por cada valor que recorro totalNotas pasa a ser =+ a ese valor
	}
	avg = totalNotas / float64(len(values)) // promedio es igual a la suma total de las notas, dividida la cantidad de valores asignadas a la función
	return avg
}

func operation(function string) (func(...float64) float64, error) {

	switch function {
	case minimum:
		return minFunc, nil

	case maximum:
		return maxFunc, nil

	case average:
		return avgFunc, nil

	default:
		msgError := fmt.Sprintf("La función %s no se encuentra definida.", function)
		return nil, errors.New(msgError)
	}
}

func main() {
	minValue, err := operation(minimum)
	if err != nil {
		fmt.Println(err)
		return
	}

	maxValue, err := operation(maximum)
	if err != nil {
		fmt.Println(err)
		return
	}

	averageValue, err := operation(average)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("La nota mínima es %.2f\n", minValue(7, 5, 3, 4, 10, 3, 4, 5))
	fmt.Printf("La nota máxima es %.2f\n", maxValue(2, 3, 7, 4, 9, 2, 4, 5))
	fmt.Printf("La nota promedio es %.2f\n", averageValue(2, 8, 6, 4, 7, 6, 4, 5))

}
