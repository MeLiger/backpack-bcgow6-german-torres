package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (s StubStore) Read(data interface{}) error {
	users := data.(*[]User)

	stubData := []User{
		{
			Id:     1,
			Name:   "Nico",
			Email:  "nico@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
		{
			Id:     2,
			Name:   "Gabi",
			Email:  "gabi@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
	}
	*users = stubData
	return nil
}

func (s StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repository := NewRepository(myStubStore)
	dataEsperada := []User{
		{
			Id:     1,
			Name:   "Nico",
			Email:  "nico@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
		{
			Id:     2,
			Name:   "Gabi",
			Email:  "gabi@mercadolibre.cl",
			Age:    1,
			Height: 1293,
			Active: true,
			Date:   "25-11-2012",
		},
	}

	//act
	resultado, _ := repository.GetAll()

	//assert
	assert.Equal(t, dataEsperada, resultado)
}
