package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	read, err := os.ReadFile("./products.csv")
	if err != nil {
		panic(err)
	}
	data := string(read)
	fmt.Println(strings.ReplaceAll(data, ";", "\t\t\t"))

}
