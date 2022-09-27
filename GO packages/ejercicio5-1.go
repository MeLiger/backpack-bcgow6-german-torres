package main

import (
	"fmt"
	"os"
)

type Product struct {
	ID       string
	Price    int
	Quantity int
}

func (p Product) details() string {
	return fmt.Sprintf("ID: %s\nPrice: %d\nQuantity: %d\n\n", p.ID, p.Price, p.Quantity)
}

func main() {
	product1 := Product{
		ID:       "KT-303",
		Price:    350,
		Quantity: 15,
	}
	product2 := Product{
		ID:       "T1-112",
		Price:    500,
		Quantity: 92,
	}

	d1 := []byte("hello")
	err := os.WriteFile("/tmp/dat1", d1, 0644)

	file, err := os.OpenFile("products.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	if _, err := f.Write([]byte(product1.details())); err != nil {
		panic(err)
	}
	if _, err := f.Write([]byte(product2.details())); err != nil {
		panic(err)
	}

}
