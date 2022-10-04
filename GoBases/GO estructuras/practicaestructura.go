package main

import "fmt"

type Persona struct {
	Nombre    string
	Genero    string
	Edad      int
	Profesión string
	Peso      float64
	Ficha     Skills
}

type Skills struct {
	Deportes    string
	Computación string
	Seguridad   string
	Ocultismo   string
}

func main() {

	skillsForNahuel := Skills{
		Deportes:  "Fútbol",
		Ocultismo: "Magia Negra",
	}

	p2 := Persona{
		Nombre:    "Nahuel",
		Genero:    "M",
		Edad:      30,
		Profesión: "Doctor",
		Peso:      77,
		Ficha:     skillsForNahuel,
	}

	p3 := new(Persona)
	p3.Nombre = "Jean Carlo"
	p3.Edad = 29
	p3.Peso = 67

	fmt.Println(p2)
	fmt.Println(p3)

}
