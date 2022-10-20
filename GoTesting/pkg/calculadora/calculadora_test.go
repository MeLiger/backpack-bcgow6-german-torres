package calculadora

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubGood(t *testing.T) {
	// Arrange (given)
	num1 := 10
	num2 := 5
	resultadoEsperado := 5

	// Act  (when)
	resultado := Sub(num1, num2)

	// Assert (then)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}

func TestSubBad(t *testing.T) {
	// Arrange (given)
	num1 := 10
	num2 := 5
	resultadoEsperado := 7

	// Act  (when)
	resultado := Sub(num1, num2)

	// Assert (then)
	assert.NotEqual(t, resultadoEsperado, resultado, "deben ser iguales")

}



//Testing emprolijado gracias a Lu Masiero.

func TestSliceSort(t *testing.T) {
	// Arrange
	scl := []int{400, 600, 100, 300, 500, 200, 900}
	expectedSlice := []int{100, 200, 300, 400, 500, 600, 900}

	// Act
	SliceSort(scl)

	// Assert
	assert.Equal(t, scl, expectedSlice, "deben ser iguales")
}


// Testing humildemente hecho por mi.

/*
func TestSliceSort(t *testing.T) {
	// Arrange
	scl := []int{400, 600, 100, 300, 500, 200, 900}
	expectedSlice := []int{100, 200, 300, 400, 500, 600, 900}

	// Act
	resultingSlice := SliceSort(scl)

	// Assert
	assert.Equal(t, resultingSlice, expectedSlice, "deben ser iguales")
} */


