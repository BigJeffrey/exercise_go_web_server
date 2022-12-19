package services

import (
	"errors"
	"net/http"
	"webserver/interfaces"
	"webserver/models"

	"github.com/google/uuid"
)

type AddProductToBasketService struct {
	Dao interfaces.PostgresqlInterface
}

func (a *AddProductToBasketService) Add(basketId uuid.UUID, productId uuid.UUID, quantity int) (int, error) {

	basketRetrieved, err := a.Dao.GetBasketById(basketId)
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}
	if basketRetrieved == nil {
		return http.StatusNotFound, errors.New("basket not found")
	}

	product, err := a.Dao.GetProductById(productId)
	if err != nil {
		return http.StatusInternalServerError, errors.New("something went wrong")
	}
	if product == nil {
		return http.StatusNotFound, errors.New("product not found")
	}

	if product.Quantity < quantity {
		return http.StatusBadRequest, errors.New("not enough products")
	}

	productInBasket, err := a.Dao.GetProductFromBasketById(basketId, productId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if productInBasket == nil {
		err := a.createProductInBasket(basketId, productId, quantity)
		if err != nil {
			return http.StatusInternalServerError, errors.New("something went wrong")
		}
	} else {
		err = a.updateProductInBasket(*productInBasket, quantity)
		if err != nil {
			return http.StatusInternalServerError, err
		}
	}

	err = a.decreaseProductQuantity(product, quantity)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (a *AddProductToBasketService) createProductInBasket(basketId uuid.UUID, productId uuid.UUID, quantity int) error {
	newBasket := models.BasketProducts{
		BasketID:  basketId,
		ProductID: productId,
		Quantity:  quantity,
	}
	_, err := a.Dao.AddProductToBasket(newBasket)

	return err

}

func (a *AddProductToBasketService) updateProductInBasket(basketProducts models.BasketProducts, quantity int) error {

	basketProducts.Quantity = basketProducts.Quantity + quantity

	return a.Dao.UpdateProductInBasket(basketProducts)

}

func (a *AddProductToBasketService) decreaseProductQuantity(product *models.Product, quantity int) error {
	product.Quantity -= quantity
	if product.Quantity == 0 {
		product.Status = "unavailable"
	} else {
		product.Status = "ordered"
	}
	return a.Dao.UpdateProduct(*product)
}
