package main

import "fmt"

const (
	dog       = "dog"
	cat       = "cat"
	tarantula = "tarantula"
	hamster   = "hamster"
)

func Animal(animals string) (func(food float64) float64, string) {
	switch animals {
	case "dog":
		return animalDog, ""
	case "cat":
		return animalCat, ""
	case "hamster":
		return animalHamster, ""
	case "tarantula":
		return animalTarantula, ""
	default:
		return nil, "No hay registros de este tipo de animal"
	}
}

func animalDog(food float64) float64 {
	return food * 10
}

func animalCat(food float64) float64 {
	return food * 5
}

func animalHamster(food float64) float64 {
	return food * 0.25
}

func animalTarantula(food float64) float64 {
	return food * 0.15
}

func main() {
	var amount float64

	amount += animalDog(5)
	amount += animalCat(7)
	amount += animalTarantula(4)
	amount += animalHamster(3)

	fmt.Printf("La cantidad de alimento necesaria es %.2fkg\n", amount)
}
