package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
