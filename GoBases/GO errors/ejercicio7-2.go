package main

import (
	"errors"
	"fmt"
	"os"
)

type MyCustomError struct {
	msg string
}

var errSalary = errors.New("el salario ingresado no alcanza el mínimo imponible")

func MyCustomErrorTest(salary int) (err error) {
	if salary < 150000 {
		err = errSalary
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
