package interfaces

import (
	"database/sql"
	"github.com/google/uuid"
	"webserver/models"
)

type PostgresqlInterface interface {
	Disconnect()
	CreateBasket() (uuid.UUID, error)
	GetBaskets() ([]models.BasketProducts, error)
	GetBasketById(id uuid.UUID) (*models.Basket, error)
	AddProductToBasket(basketSum models.BasketProducts) (uuid.UUID, error)
	GetProductFromBasketById(basketId uuid.UUID, productId uuid.UUID) (*models.BasketProducts, error)
	UpdateProductInBasket(basketSum models.BasketProducts) error
	DeleteProductFromBasket(basketSum models.BasketProducts) (sql.Result, error)
	CreateProduct(models.Product) (uuid.UUID, error)
	GetProducts() ([]models.Product, error)
	GetProductById(id uuid.UUID) (*models.Product, error)
	UpdateProduct(models.Product) error
	DeleteProduct(id uuid.UUID) (sql.Result, error)
	FilterProduct(search string) ([]models.Product, error)
	GetProductsByPrice(priceSelect models.PriceSelectRequest) ([]models.Product, error)
}
