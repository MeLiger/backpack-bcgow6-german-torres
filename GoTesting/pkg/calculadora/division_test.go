package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	// Arrange (given)
	num := 10
	den := 0
	expectedResult := errorDen

	// Act  (when)
	resultado, err := Div(num, den)

	// Assert (then)
	assert.Nil(t, err)
	assert.Equal(t, resultado, expectedResult, "deben ser iguales")

}
