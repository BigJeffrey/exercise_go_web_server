package mocks

import (
	"database/sql"
	"fmt"
	"webserver/models"

	"github.com/stretchr/testify/mock"

	"github.com/google/uuid"
)

type PostgresqlMock struct {
	mock.Mock
}

func (m *PostgresqlMock) CreateProduct(product models.Product) (uuid.UUID, error) {
	//to implement
	return product.ID, nil
}

func (m *PostgresqlMock) GetProducts() ([]models.Product, error) {
	//to implement
	return nil, nil

}

func (m *PostgresqlMock) GetProductById(id uuid.UUID) (*models.Product, error) {
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
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.GetProductById(%d) failed because object wasn't correct type: %v", id, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.GetProductById(%d) failed because object wasn't correct type: %v", id, args.Get(1)))
	}

	return nil, nil

}

func (m *PostgresqlMock) UpdateProduct(product models.Product) error {
	args := m.Called(product)

	returnError, errorOK := args.Get(0).(error)
	if errorOK {
		return returnError
	}

	if args.Get(0) == nil {
		return nil
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.UpdateProduct(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	return nil
}

func (m *PostgresqlMock) DeleteProduct(uuid.UUID) (sql.Result, error) {
	//to implement
	return nil, nil

}

func (m *PostgresqlMock) FilterProduct(string) ([]models.Product, error) {
	//to implement
	return nil, nil

}

func (m *PostgresqlMock) CreateBasket() (uuid.UUID, error) {
	//to implement
	return [16]byte{}, nil

}

func (m *PostgresqlMock) GetProductsByPrice(models.PriceSelectRequest) ([]models.Product, error) {
	//to implement
	return nil, nil

}

func (m *PostgresqlMock) GetBaskets() ([]models.BasketProducts, error) {
	//to implement
	return nil, nil

}

func (m *PostgresqlMock) GetBasketById(id uuid.UUID) (*models.Basket, error) {
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
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.GetBasketById(%d) failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.GetBasketById(%d) failed because object wasn't correct type: %v", 0, args.Get(1)))
	}

	return nil, nil

}

func (m *PostgresqlMock) AddProductToBasket(basketProducts models.BasketProducts) (uuid.UUID, error) {
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
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.AddProductToBasket(%d) failed because object wasn't correct type: %v", returnID, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.AddProductToBasket(%d) failed because object wasn't correct type: %v", returnError, args.Get(1)))
	}

	return uuid.UUID{}, nil

}

func (m *PostgresqlMock) GetProductFromBasketById(uuid.UUID, uuid.UUID) (*models.BasketProducts, error) {
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
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.GetProductFromBasketById(%d) failed because object wasn't correct type: %v", returnBasketProducts, args.Get(0)))
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.GetProductFromBasketById(%d) failed because object wasn't correct type: %v", returnError, args.Get(1)))
	}

	return nil, nil
}

func (m *PostgresqlMock) UpdateProductInBasket(basketProducts models.BasketProducts) error {
	args := m.Called(basketProducts)

	returnError, errorOK := args.Get(0).(error)
	if errorOK {
		return returnError
	}

	if args.Get(0) == nil {
		return nil
	}

	if !errorOK {
		panic(fmt.Sprintf("assert: arguments: PostgresqlMock.UpdateProductInBasket(%d) failed because object wasn't correct type: %v", returnError, args.Get(0)))
	}

	return nil
}

func (m *PostgresqlMock) DeleteProductFromBasket(models.BasketProducts) (sql.Result, error) {
	//to implement
	return nil, nil
}

func (m *PostgresqlMock) Disconnect() {
	//to implement
}
