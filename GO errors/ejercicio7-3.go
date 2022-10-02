package main

import (
	"fmt"
	"os"
)

type MyCustomError struct {
	msg string
}

func MyCustomErrorTest(salary int) (err error) {
	if salary < 150000 {
		err = fmt.Errorf("error: el mínimo imponible es de 150000 y el salario ingresado es de: %d", salary)
		return
	}
	return
}

func main() {

	salary := 140000

	err := MyCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuestos")
}
