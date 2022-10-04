package main

import (
	"fmt"
)

type client struct {
	Legajo         int
	NombreApellido string
	DNI            int
	Telefono       int
	Domicilio      string
}

func (c client) ID() {
	fmt.Printf("Legajo: %d\n", c.Legajo)

}

func main() {

	c1 := client{19889, "Jorge García", 20988402, 46729888, "Tapalque 1930"}
	client.ID(c1)

}
