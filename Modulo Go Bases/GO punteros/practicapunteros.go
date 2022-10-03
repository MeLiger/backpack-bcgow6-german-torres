package main

import "fmt"

func main() {

	//var puntero *int   //asi se define puntero
	//p2 := new(int)  //otra forma

	var numero int = 10

	p3 := &numero //obtiene direccion de variable

	fmt.Printf("La direccion es: %p\n", &numero)

	Incrementar(p3)

	fmt.Printf("El  valor de numero es %d\n", numero)
}

func Incrementar(puntero *int) {

	*puntero = 20
}
