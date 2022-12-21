package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	postgresqldao "webserver/mocks"
	"webserver/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddProductToBasketService_Add_BasketThatNotExist(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	daoMock.On("GetBasketById", mock.Anything, mock.Anything).Return(nil, nil).Once()

	code, err := service.Add(uuid.New(), uuid.New(), 2)
	daoMock.AssertExpectations(t)

	assert.Equal(t, err.Error(), "basket not found")
	assert.Equal(t, code, http.StatusNotFound)
}

func TestAddProductToBasketService_Add_FetchBasketError(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	daoMock.On("GetBasketById", mock.Anything, mock.Anything).Return(nil, errors.New("postgresql error")).Once()

	code, err := service.Add(uuid.New(), uuid.New(), 2)
	daoMock.AssertExpectations(t)

	assert.NotNil(t, err, "add should return error")
	assert.Equal(t, err.Error(), "something went wrong")
	assert.Equal(t, code, http.StatusInternalServerError)
}

func TestAddProductToBasketService_Add_ProductThatNotExist(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.UUID{},
	}

	daoMock.On("GetBasketById", mock.Anything).Return(basket, nil).Once()
	daoMock.On("GetProductById", mock.Anything).Return(nil, nil).Once()

	code, err := service.Add(uuid.New(), uuid.New(), 2)
	daoMock.AssertExpectations(t)

	assert.Equal(t, err.Error(), "product not found")
	assert.Equal(t, code, http.StatusNotFound)
}

func TestAddProductToBasketService_Add_FetchProductError(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.UUID{},
	}

	daoMock.On("GetBasketById", mock.Anything).Return(basket, nil).Once()
	daoMock.On("GetProductById", mock.Anything, mock.Anything).Return(nil, errors.New("postgresql error")).Once()

	code, err := service.Add(uuid.New(), uuid.New(), 2)
	daoMock.AssertExpectations(t)

	assert.NotNil(t, err, "add should return error")
	assert.Equal(t, err.Error(), "something went wrong")
	assert.Equal(t, code, http.StatusInternalServerError)
}

func TestAddProductToBasketService_Add_AddMoreProductsThanInStock(t *testing.T) {

	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.New(),
	}

	product := &models.Product{
		ID:       uuid.New(),
		Quantity: 5, // starting stock quantity!
		Status:   "",
	}

	daoMock.On("GetBasketById", basket.ID).Return(basket, nil).Once()
	daoMock.On("GetProductById", product.ID).Return(product, nil).Once()

	code, err := service.Add(basket.ID, product.ID, product.Quantity+1) // more than in stock!
	daoMock.AssertExpectations(t)

	assert.NotNil(t, err, "add should return error")
	assert.Equal(t, err.Error(), "not enough products")
	assert.Equal(t, code, http.StatusBadRequest)
}

func TestAddProductToBasketService_Add_AddProductThatNotExitsInBasket(t *testing.T) {

	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.New(),
	}

	product := &models.Product{
		ID:       uuid.New(),
		Quantity: 5, // starting stock quantity!
		Status:   "",
	}

	daoMock.On("GetBasketById", basket.ID).Return(basket, nil).Once()
	daoMock.On("GetProductById", product.ID).Return(product, nil).Once()
	daoMock.On("GetProductFromBasketById").Return(nil, nil).Once()

	daoMock.On("AddProductToBasket", models.BasketProducts{
		BasketID:  basket.ID,
		ProductID: product.ID,
		Quantity:  product.Quantity,
	}).Return(uuid.New(), nil).Once()

	daoMock.On("UpdateProduct", models.Product{
		ID:       product.ID,
		Quantity: 0, // end stock quantity (-5)!
		Status:   "unavailable",
	}).Return(nil).Once()

	_, err := service.Add(basket.ID, product.ID, product.Quantity)
	daoMock.AssertExpectations(t)

	assert.True(t, daoMock.AssertNotCalled(t, "UpdateProductInBasket"))
	assert.Nil(t, err, "add should not return error")
}

func TestAddProductToBasketService_Add_AddProductThatExistsInBasket(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.New(),
	}

	product := &models.Product{
		ID:       uuid.New(),
		Quantity: 5, // starting stock quantity!
	}

	basketSum := &models.BasketProducts{
		ID:        uuid.New(),
		BasketID:  basket.ID,
		ProductID: product.ID,
		Quantity:  2, // starting basket quantity!
	}

	daoMock.On("GetBasketById", mock.Anything).Return(basket, nil).Once()
	daoMock.On("GetProductById", mock.Anything).Return(product, nil).Once()
	daoMock.On("GetProductFromBasketById").Return(basketSum, nil).Once()

	daoMock.On("UpdateProductInBasket", models.BasketProducts{
		ID:        basketSum.ID,
		BasketID:  basket.ID,
		ProductID: product.ID,
		Quantity:  4, // end basket quantity (+2)!
	}).Return(nil, nil).Once()

	daoMock.On("UpdateProduct", models.Product{
		ID:       product.ID,
		Quantity: 3, // end stock quantity (-2)!
		Status:   "ordered",
	}).Return(nil).Once()

	code, err := service.Add(basket.ID, product.ID, 2)
	daoMock.AssertExpectations(t)

	assert.True(t, daoMock.AssertNotCalled(t, "AddProductToBasket"))
	assert.Nil(t, err, "add should not return error")
	assert.Equal(t, code, http.StatusOK)
}

func TestAddProductToBasketService_ErrorFetchingProductFromBasket(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.UUID{},
	}

	product := &models.Product{
		ID:       uuid.UUID{},
		Quantity: 5,
	}

	daoMock.On("GetBasketById", mock.Anything).Return(basket, nil).Once()
	daoMock.On("GetProductById", mock.Anything).Return(product, nil).Once()
	daoMock.On("GetProductFromBasketById").Return(nil, errors.New("postgresql error")).Once()

	code, err := service.Add(basket.ID, product.ID, 2)
	daoMock.AssertExpectations(t)

	assert.NotNil(t, err, "add should return error")
	assert.Equal(t, err.Error(), "postgresql error")
	assert.Equal(t, code, http.StatusInternalServerError)
}

func TestAddProductToBasketService_ErrorAddingProductToBasket(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.UUID{},
	}

	product := &models.Product{
		ID:       uuid.UUID{},
		Quantity: 5,
	}

	daoMock.On("GetBasketById", mock.Anything).Return(basket, nil).Once()
	daoMock.On("GetProductById", mock.Anything).Return(product, nil).Once()
	daoMock.On("GetProductFromBasketById").Return(nil, errors.New("postgresql error")).Once()

	code, err := service.Add(basket.ID, product.ID, 2)
	daoMock.AssertExpectations(t)

	assert.NotNil(t, err, "add should return error")
	assert.Equal(t, err.Error(), "postgresql error")
	assert.Equal(t, code, http.StatusInternalServerError)
}

func TestAddProductToBasketService_ErrorUpdatingProductInBasket(t *testing.T) {
	daoMock := &postgresqldao.PostgresqlMock{}

	service := AddProductToBasketService{
		Dao: daoMock,
	}

	basket := &models.Basket{
		ID: uuid.UUID{},
	}

	product := &models.Product{
		ID:       uuid.UUID{},
		Quantity: 5,
	}

	basketSum := &models.BasketProducts{}

	daoMock.On("GetBasketById", mock.Anything).Return(basket, nil).Once()
	daoMock.On("GetProductById", mock.Anything).Return(product, nil).Once()
	daoMock.On("GetProductFromBasketById").Return(basketSum, nil).Once()
	daoMock.On("UpdateProductInBasket", mock.Anything).Return(errors.New("postgresql error")).Once()

	code, err := service.Add(basket.ID, product.ID, 2)
	daoMock.AssertExpectations(t)

	assert.NotNil(t, err, "add should return error")
	assert.Equal(t, err.Error(), "postgresql error")
	assert.Equal(t, code, http.StatusInternalServerError)
}
