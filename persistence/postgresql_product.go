package persistence

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"webserver/models"
)

func (m *Postgresql) GetProducts() ([]models.Product, error) {
	var product models.Product
	var products []models.Product
	sqlStatement := `SELECT * FROM products`

	rows, err := m.client.Query(sqlStatement)
	if err != nil {
		return nil, err
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

func (m *Postgresql) GetProductById(id uuid.UUID) (*models.Product, error) {
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

func (m *Postgresql) GetProductsByPrice(priceSelect models.PriceSelectRequest) ([]models.Product, error) {
	var product models.Product
	var products []models.Product

	var sqlStatement string
	switch priceSelect.PriceCondition {
	case "greater":
		sqlStatement = `SELECT * FROM products WHERE price>$1`
	case "equal":
		sqlStatement = `SELECT * FROM products WHERE price=$1`
	case "lower":
		sqlStatement = `SELECT * FROM products WHERE price<$1`
	default:
		return nil, errors.New("unrecognized condition")
	}

	rows, err := m.client.Query(sqlStatement, priceSelect.Price)
	if err != nil {
		return nil, err
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

func (m *Postgresql) CreateProduct(product models.Product) (uuid.UUID, error) {
	var id uuid.UUID
	sqlStatement := `INSERT INTO products (name, description, price, quantity, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := m.client.QueryRow(sqlStatement, product.Name, product.Description, product.Price, product.Quantity, product.Status).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (m *Postgresql) UpdateProduct(product models.Product) error {
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

func (m *Postgresql) DeleteProduct(id uuid.UUID) (sql.Result, error) {
	sqlStatement := `DELETE FROM products WHERE id=$1`
	return m.client.Exec(sqlStatement, id)
}

func (m *Postgresql) FilterProduct(searche string) ([]models.Product, error) {
	var product models.Product
	var products []models.Product

	sqlStatement := `SELECT * FROM products WHERE name LIKE $1 OR description LIKE $1 OR status LIKE $1`

	rows, err := m.client.Query(sqlStatement, "%"+searche+"%")
	if err != nil {
		return nil, err
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
