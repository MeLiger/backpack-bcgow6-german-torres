package main

import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	employees["Federico"] = 25
	delete(employees, "Pedro")
	fmt.Printf("La edad de Benjamin es: %v\n", employees["Benjamin"])
	for k, v := range employees {
		if v > 21 {
			fmt.Println(k, "=>", v)
		}
	}

}
