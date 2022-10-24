package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	ReadWasCalled bool
}

func (s StubStore) Read(data interface{}) error {
	s.ReadWasCalled = true

	us := data.(*[]User)
	stubData := []User{
		{
			Id:     1,
			Name:   "User1",
			Email:  "user1@mercadolibre.com.ar",
			Age:    22,
			Height: 178,
			Active: true,
			Date:   "10-11-2006",
		},
		{
			Id:     2,
			Name:   "User2",
			Email:  "user2@mercadolibre.com.co",
			Age:    28,
			Height: 169,
			Active: true,
			Date:   "09-07-2014",
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
			Age:    22,
			Height: 178,
			Active: true,
			Date:   "10-11-2006",
		},
		{
			Id:     2,
			Name:   "User2",
			Email:  "user2@mercadolibre.com.co",
			Age:    28,
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

func TestUpdateName(t *testing.T) {
	//arrange
	myStubStore := StubStore{}
	repository := NewRepository(&myStubStore)
	expectedResult := User{
		Id:     1,
		Name:   "User1 After Update",
		Email:  "user1@mercadolibre.com.ar",
		Age:    24,
		Height: 178,
		Active: true,
		Date:   "10-11-2006",
	}

	/*	initialData := []User{
			{
				Id:     1,
				Name:   "User1 Before Update",
				Email:  "user1@mercadolibre.com.ar",
				Age:    22,
				Height: 178,
				Active: true,
				Date:   "10-11-2006",
			},
		}

		store := MockStorage{
			mockedData:    initialData,
			readWasCalled: false,
		}
	*/

	result, err := repository.UpdateNameAge(1, "User1 After Update", 24)

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result, "different result than expected")
	assert.True(t, myStubStore.ReadWasCalled, "error on Read function")

}
