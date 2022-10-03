package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.ReadFile("./panicpractic.go")
	if err != nil {
		panic(err)

	}

	var slice []int

	s := slice[4]
	fmt.Println(s)
}
