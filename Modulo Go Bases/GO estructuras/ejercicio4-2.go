package main

import (
	"fmt"
	"math"
)

type Matrix struct {
	Values    []float64
	Height    int
	Width     int
	Quadratic bool
	MaxValue  float64
}

func (m Matrix) isQuadratic() bool {
	if (m.Height == m.Width) && m.Height != 0 {
		return true
	}
	return false
}

func (m Matrix) Set(values ...float64) {
	m.Values = values

	if m.Quadratic == false {
		fmt.Println("Es Cuadratica")
	}
}

func (m Matrix) Print() {
	for i := 0; i < m.Height; i++ {
		for j := 0; j < m.Width; j++ {
			fmt.Printf("%f\t", m.Values[i*m.Width+j])
		}
		fmt.Printf("\n")
	}
}

func (m Matrix) Max() float64 {
	max := -math.MaxFloat64
	for _, elem := range m.Values {
		if elem > max {
			m.MaxValue = elem
		}
	}
	return max
}
func main() {

	m := Matrix{Height: 3, Width: 3, Quadratic: true}
	m.Set(9, 8, 7, 6, 5, 4, 3, 2, 1)
	m.Print()

}

/*s1 := [][]int{
	{1, 2},
	{3, 4},
	{5, 6},
	{7, 8},
}

Accessing multi-dimensional slice
fmt.Println("Slice 1 : ", s1)

// Creating multi-dimensional slice
s2 := [][]string{
	[]string{"str1", "str2"},
	[]string{"str3", "str4"},
	[]string{"str5", "str6"},
}

// Accessing multi-dimensional slice
fmt.Println("Slice 2 : ", s2)

// Printing every slice inside s2
fmt.Println("MultiDimensional Slice s2:")
for i := 0; i < len(s2); i++ {
	fmt.Println(s2[i])
}

// Printing elements in slice as matrix
fmt.Println("Slice s2 Like Matrix:")

// number of rows in s2
n := len(s2)

// number of columns in s2
m := len(s2[0])
for i := 0; i < n; i++ {
	for j := 0; j < m; j++ {
		fmt.Print(s2[i][j] + " ")
	}
	fmt.Print("\n")
}
*/
