package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Ejecución Finalizada")
	}()
	_, err := os.Open("customer.txt")
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}

}
