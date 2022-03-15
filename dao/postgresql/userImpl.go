package postgresqldao

import (
	"database/sql"
	"errors"
	"fmt"
	"webserver/models"

	"github.com/google/uuid"
)

func (m *PostgreSql) CreateProduct(product models.Product) (uuid.UUID, error) {
	var id uuid.UUID
	sqlStatement := `INSERT INTO products (name, description, price, quantity, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := m.client.QueryRow(sqlStatement, product.Name, product.Description, product.Price, product.Quantity, product.Status).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *PostgreSql) GetProducts() ([]models.Product, error) {
	var product models.Product
	var products []models.Product
	sqlStatement := `SELECT * FROM products`

	rows, err := m.client.Query(sqlStatement)
	if err != nil {
		return []models.Product{}, err
	}

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Status)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (m *PostgreSql) GetProductById(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	sqlStatement := `SELECT * FROM products WHERE id=$1`
	err := m.client.QueryRow(sqlStatement, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Status)

	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return &product, nil
}

func (m *PostgreSql) UpdateProduct(product models.Product) error {
	sqlStatement := `UPDATE products SET name=$1, description=$2, price=$3, quantity=$4, status=$5 WHERE id=$6`
	result, err := m.client.Exec(sqlStatement, product.Name, product.Description, product.Price, product.Quantity, product.Status, product.ID)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		strErr := fmt.Sprintf("Rows affected: %d /updateproduct", count)
		return errors.New(strErr)
	}

	return nil
}

func (m *PostgreSql) DeleteProduct(id uuid.UUID) (sql.Result, error) {
	sqlStatement := `DELETE FROM products WHERE id=$1`
	return m.client.Exec(sqlStatement, id)
}

func (m *PostgreSql) FilterProduct(searche string) (*[]models.Product, error) {
	var product models.Product
	var products []models.Product

	sqlStatement := `SELECT * FROM products WHERE name LIKE $1 OR description LIKE $1 OR status LIKE $1`

	rows, err := m.client.Query(sqlStatement, "%"+searche+"%")
	if err != nil {
		return &[]models.Product{}, err
	}

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Status)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return &products, nil
}

func (m *PostgreSql) CreateBasket() (uuid.UUID, error) {
	var id uuid.UUID
	sqlStatement := `INSERT INTO baskets DEFAULT VALUES RETURNING id`
	err := m.client.QueryRow(sqlStatement).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *PostgreSql) GetProductsByPrice(priceSelect models.PriceSelectRequest) ([]models.Product, error) {
	var product models.Product
	var products []models.Product

	var sqlStatement string
	if priceSelect.PriceCondition == "greater" {
		sqlStatement = `SELECT * FROM products WHERE price>$1`
	} else if priceSelect.PriceCondition == "equal" {
		sqlStatement = `SELECT * FROM products WHERE price=$1`
	} else if priceSelect.PriceCondition == "lower" {
		sqlStatement = `SELECT * FROM products WHERE price<$1`
	} else {
		return []models.Product{}, errors.New("unrecognized condition")
	}

	rows, err := m.client.Query(sqlStatement, priceSelect.Price)
	if err != nil {
		return []models.Product{}, err
	}

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Status)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (m *PostgreSql) GetBaskets() ([]models.BasketProducts, error) {
	var basket models.BasketProducts
	var baskets []models.BasketProducts
	sqlStatement := `SELECT * FROM basket_products`

	rows, err := m.client.Query(sqlStatement)
	if err != nil {
		return []models.BasketProducts{}, err
	}

	for rows.Next() {
		err := rows.Scan(&basket.ID, &basket.BasketID, &basket.ProductID, &basket.Quantity)
		if err != nil {
			return nil, err
		}
		baskets = append(baskets, basket)
	}
	return baskets, nil
}

func (m *PostgreSql) GetBasketById(id uuid.UUID) (*models.Basket, error) {
	var basket models.Basket
	sqlStatement := `SELECT * FROM baskets WHERE id=$1`
	err := m.client.QueryRow(sqlStatement, id).Scan(&basket.ID)

	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		} else {
			return nil, nil
		}
	}

	return &basket, nil
}

func (m *PostgreSql) AddProductToBasket(basketProducts models.BasketProducts) (uuid.UUID, error) {
	var id uuid.UUID
	sqlStatement := `INSERT INTO basket_products (basketid, productid, quantity) VALUES ($1, $2, $3) RETURNING id`
	err := m.client.QueryRow(sqlStatement, basketProducts.BasketID, basketProducts.ProductID, basketProducts.Quantity).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *PostgreSql) GetProductFromBasketById(basketId uuid.UUID, productId uuid.UUID) (*models.BasketProducts, error) {
	var basketProducts models.BasketProducts
	sqlStatement := `SELECT * FROM basket_products WHERE productid=$1 AND basketid=$2`
	err := m.client.QueryRow(sqlStatement, basketId, productId).Scan(&basketProducts.ID, &basketProducts.BasketID, &basketProducts.ProductID, &basketProducts.Quantity)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return &models.BasketProducts{}, err
		} else {
			return nil, nil
		}
	}

	return &basketProducts, nil
}

func (m *PostgreSql) UpdateProductInBasket(basketProducts models.BasketProducts) error {
	sqlStatement := `UPDATE basket_products SET basketid=$1, productid=$2, quantity=$3 WHERE id=$4`

	result, err := m.client.Exec(sqlStatement, basketProducts.BasketID, basketProducts.ProductID, basketProducts.Quantity, basketProducts.ID)
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		strErr := fmt.Sprintf("Rows affected: %d /updateProductInBasket", count)
		return errors.New(strErr)
	}

	return nil
}

func (m *PostgreSql) DeleteProductFromBasket(basketProducts models.BasketProducts) (sql.Result, error) {
	sqlStatement := `DELETE FROM basket_products WHERE basketid=$1 AND productid=$2`
	return m.client.Exec(sqlStatement, basketProducts.BasketID, basketProducts.ProductID)
}
