package main

import "fmt"

type Products struct {
	Name     string
	Price    float64
	Quantity int
}

type Services struct {
	Name    string
	Price   float64
	Minutes int
}

type Upkeep struct {
	Name  string
	Price float64
}

func main() {

	products := []Products{
		{Name: "Lavandina",
			Price:    350.50,
			Quantity: 20,
		},
		{
			Name:     "Jabón",
			Price:    200,
			Quantity: 5,
		},
	}

	services := []Services{
		{Name: "Lavado",
			Price:   300,
			Minutes: 180,
		},

		{Name: "Planchado",
			Price:   200,
			Minutes: 180,
		},

		{Name: "Limpieza general",
			Price:   300,
			Minutes: 180,
		},
	}

	upkeep := []Upkeep{
		{
			Name:  "Lavarropa",
			Price: 2000,
		},
		{
			Name:  "Fotocopiadora",
			Price: 3000,
		},
		{
			Name:  "Computadora",
			Price: 1500,
		},
	}
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	go addProducts(&products, c1)
	go addServices(&services, c2)
	go addUpkeep(&upkeep, c3)

	total1 := <-c1
	total2 := <-c2
	total3 := <-c3

	fmt.Println("* Total final * $ ", total1+total2+total3)

}

func addProducts(products *[]Products, c chan float64) {
	var total float64
	for _, value := range *products {
		total += value.Price * float64(value.Quantity)
	}

	fmt.Println("Precio total de productos: $", total)
	c <- total
	close(c)
}

func addServices(services *[]Services, c chan float64) {
	var total float64
	for _, value := range *services {
		if value.Minutes < 30 {
			total += value.Price
		} else {
			total += value.Price * float64(value.Minutes) / 30
		}
		fmt.Println("Precio total de servicio: $", total)
		c <- total
		close(c)

	}

}
func addUpkeep(upkeep *[]Upkeep, c chan float64) {
	var total float64
	for _, value := range *upkeep {
		total += value.Price
	}
	fmt.Println("Precio total de mantenimiento: $", total)
	c <- total
	close(c)
}
