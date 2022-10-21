package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (s StubStore) Read(data interface{}) error {
	us := data.(*[]User)
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
	*us = stubData
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
			Name:   "User1",
			Email:  "user1@mercadolibre.com.ar",
			Age:    1,
			Height: 178,
			Active: true,
			Date:   "10-11-2006",
		},
		{
			Id:     2,
			Name:   "User2",
			Email:  "user2@mercadolibre.com.co",
			Age:    1,
			Height: 169,
			Active: true,
			Date:   "09-07-2014",
		},
	}

	//act
	resultado, _ := repository.GetAll()

	//assert
	assert.Equal(t, dataEsperada, resultado)
}

/*func TestUpdateName(t *testing.T) {
//arrange
myStubStore := StubStore{}
repository := NewRepository(&myStubStore)
expectedResult := User{

//act
result := repository.Update

//assert
assert.Equal(t, expectedResult, result)
*/
