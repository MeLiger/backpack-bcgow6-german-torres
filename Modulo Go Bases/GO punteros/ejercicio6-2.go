package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Users struct {
	Name     string
	Surname  string
	Email    string
	Products []Product
}

func newProduct(name string, price float64, quantity int, product *Product) {
	(*product).Name = name
	(*product).Price = price
	(*product).Quantity = quantity

}

func addProduct(user *Users, product *Product, quantity int) {
	product.Quantity *= quantity
	user.Products = append(user.Products, *product)
}

func deleteProduct(user *Users) {
	user.Products = []Product{}
}

func main() {

	p1 := Product{"reloj", 2500, 20}
	p2 := Product{"billetera", 350, 35}
	p3 := Product{"mochila", 950, 10}
	productList := []Product{p1, p2, p3}

	u1 := Users{
		Name:    "Elliot",
		Surname: "Glee",
		Email:   "elliotpez@gmail.com",
	}

	u2 := Users{
		Name:    "Charles",
		Surname: "Nixon",
		Email:   "charlienix@gmail.com",
	}

	fmt.Println(u1, "\n", u2)
	fmt.Println(productList)
}
