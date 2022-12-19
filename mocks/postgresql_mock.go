package mocks

import (
	"database/sql"
	"fmt"
	"webserver/models"

	"github.com/stretchr/testify/mock"

	"github.com/google/uuid"
)

type PostgreSqlMock struct {
	mock.Mock
}

func (m *PostgreSqlMock) CreateProduct(product models.Product) (uuid.UUID, error) {
	//to implement
	return product.ID, nil
}

func (m *PostgreSqlMock) GetProducts() ([]models.Product, error) {
	//to implement
	return nil, nil

}

func (m *PostgreSqlMock) GetProductById(id uuid.UUID) (*models.Product, error) {
	args := m.Called(id)

	returnProduct, productOK := args.Get(0).(*models.Product)
	if productOK {
		return returnProduct, nil
	}

	returnError, errorOK := args.Get(1).(error)
	if errorOK {
		return nil, returnError
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	if !productOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.GetProductById(%d) failed because object wasn't correct type: %v", id, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.GetProductById(%d) failed because object wasn't correct type: %v", id, args.Get(1)))
	}

	return nil, nil

}

func (m *PostgreSqlMock) UpdateProduct(product models.Product) error {
	args := m.Called(product)

	returnError, errorOK := args.Get(0).(error)
	if errorOK {
		return returnError
	}

	if args.Get(0) == nil {
		return nil
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.UpdateProduct(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	return nil
}

func (m *PostgreSqlMock) DeleteProduct(uuid.UUID) (sql.Result, error) {
	//to implement
	return nil, nil

}

func (m *PostgreSqlMock) FilterProduct(string) ([]models.Product, error) {
	//to implement
	return nil, nil

}

func (m *PostgreSqlMock) CreateBasket() (uuid.UUID, error) {
	//to implement
	return [16]byte{}, nil

}

func (m *PostgreSqlMock) GetProductsByPrice(models.PriceSelectRequest) ([]models.Product, error) {
	//to implement
	return nil, nil

}

func (m *PostgreSqlMock) GetBaskets() ([]models.BasketProducts, error) {
	//to implement
	return nil, nil

}

func (m *PostgreSqlMock) GetBasketById(id uuid.UUID) (*models.Basket, error) {
	args := m.Called(id)

	returnBasket, basketOK := args.Get(0).(*models.Basket)
	if basketOK {
		return returnBasket, nil
	}

	returnError, errorOK := args.Get(1).(error)
	if errorOK {
		return nil, returnError
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	if !basketOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.GetBasketById(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.GetBasketById(%d) failed because object wasn't correct type: %v", 0, args.Get(1)))
	}

	return nil, nil

}

func (m *PostgreSqlMock) AddProductToBasket(basketProducts models.BasketProducts) (uuid.UUID, error) {
	args := m.Called(basketProducts)

	returnID, idOK := args.Get(0).(uuid.UUID)
	if idOK {
		return returnID, nil
	}

	returnError, errorOK := args.Get(1).(error)
	if errorOK {
		return uuid.UUID{}, returnError
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return uuid.UUID{}, nil
	}

	if !idOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.AddProductToBasket(%d) failed because object wasn't correct type: %v", returnID, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.AddProductToBasket(%d) failed because object wasn't correct type: %v", returnError, args.Get(1)))
	}

	return uuid.UUID{}, nil

}

func (m *PostgreSqlMock) GetProductFromBasketById(uuid.UUID, uuid.UUID) (*models.BasketProducts, error) {
	args := m.Called()

	returnBasketProducts, basketProductsOK := args.Get(0).(*models.BasketProducts)
	if basketProductsOK {
		return returnBasketProducts, nil
	}

	returnError, errorOK := args.Get(1).(error)
	if errorOK {
		return nil, returnError
	}

	if args.Get(0) == nil && args.Get(1) == nil {
		return nil, nil
	}

	if !basketProductsOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.GetProductFromBasketById(%d) failed because object wasn't correct type: %v", returnBasketProducts, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.GetProductFromBasketById(%d) failed because object wasn't correct type: %v", returnError, args.Get(1)))
	}

	return nil, nil
}

func (m *PostgreSqlMock) UpdateProductInBasket(basketProducts models.BasketProducts) error {
	args := m.Called(basketProducts)

	returnError, errorOK := args.Get(0).(error)
	if errorOK {
		return returnError
	}

	if args.Get(0) == nil {
		return nil
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgreSqlMock.UpdateProductInBasket(%d) failed because object wasn't correct type: %v", returnError, args.Get(0)))
	}

	return nil
}

func (m *PostgreSqlMock) DeleteProductFromBasket(models.BasketProducts) (sql.Result, error) {
	//to implement
	return nil, nil
}

func (m *PostgreSqlMock) Disconnect() {
	//to implement
}
