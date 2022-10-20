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
