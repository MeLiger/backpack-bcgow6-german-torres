package main

import (
	"fmt"
	"os"
)

type MyCustomError struct {
	msg string
}

func (err *MyCustomError) Error() string {
	return err.msg
}

func MyCustomErrorTest(salary int) (err error) {
	if salary < 150000 {
		err = &MyCustomError{"error: el salario ingresado no alcanza el mínimo imponible"}
		return
	}
	return
}

func main() {

	salary := 180000

	err := MyCustomErrorTest(salary)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Debe pagar impuestos")
}
