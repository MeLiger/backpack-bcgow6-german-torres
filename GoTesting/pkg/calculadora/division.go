package calculadora

import "errors"

var errorDen = errors.New("denominator cant be 0")

// Función que recibe numerador y denominador y retorna la division resultante.
func Div(num, den int) (int, error) {
	if den == 0 {
		return 0, errorDen
	}

	return num / den, nil
}
