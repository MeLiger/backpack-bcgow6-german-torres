package calculadora

import (
	"sort"
)

// Función que recibe dos enteros y retorna la suma resultante
func Add(num1, num2 int) int {
	return num1 + num2
}

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Sub(num1, num2 int) int {
	return num1 - num2
}


//Funcion que hice humildemente yo.

/*   
func SliceSort(scl []int) []int {
	sort.Ints(scl)
	return scl
}  */

//Función emprolijada gracias a Lu Masiero.

func SliceSort(scl []int) {
	sort.Ints(scl)
	
}
