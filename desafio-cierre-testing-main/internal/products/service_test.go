package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepository struct {
	stubProd    []Product
	stubProdErr error
}

func (sr *StubRepository) GetAllBySeller(sellerID string) ([]Product, error) {
	if sr.stubProdErr != nil {
		return []Product{}, sr.stubProdErr
	}
	return sr.stubProd, nil
}

func TestGetAllbySeller(t *testing.T) {
	//arrange
	expectedProduct := []Product{{ID: "1", SellerID: "111", Description: "desc", Price: 100}}

	myStubRepo := StubRepository{stubProd: expectedProduct}
	service := NewService(&myStubRepo)

	//act
	result, err := service.GetAllBySeller("111")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expectedProduct, result)

}

func TestGetAllbySellerFail(t *testing.T) {
	//arrange
	expectedProduct := []Product{{ID: "1", SellerID: "111", Description: "desc", Price: 100}}

	myStubRepo := StubRepository{stubProd: expectedProduct}
	service := NewService(&myStubRepo)

	//act
	result, err := service.GetAllBySeller("111")

	//assert
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestIntegration(t *testing.T) {
	//arrange
	repo := NewRepository()
	service := NewService(repo)
	expected := []Product{{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55}}

	//act
	result, err := service.GetAllBySeller("FEX112AC")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
