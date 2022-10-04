/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type Student struct {
	Name    string
	Surname string
	Dni     string
	Date    string
}

func (s Student) detalle() {
	fmt.Printf("Name: %s\n", s.Name)
	fmt.Printf("Surname: %s\n", s.Surname)
	fmt.Printf("Dni: %s\n", s.Dni)
	fmt.Printf("Date: %s\n", s.Date)
	fmt.Println("")

}

func main() {

	s1 := Student{"Eliseo", "Molina", "33103987", "26-9-22"}
	Student.detalle(s1)

	s2 := Student{
		Name:    "Johnatan",
		Surname: "Figueroa",
		Dni:     "29889012",
		Date:    "10-9-22",
	}
	Student.detalle(s2)

}
