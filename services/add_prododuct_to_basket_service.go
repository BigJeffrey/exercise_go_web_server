package services

import (
	"errors"
	"net/http"
	"webserver/dao/interfaces"
	"webserver/models"

	"github.com/google/uuid"
)

type AddProductToBasketService struct {
	Dao interfaces.AppDao
}

func (a *AddProductToBasketService) Add(basketId uuid.UUID, productId uuid.UUID, quantity int) (error, int) {

	basketRetrieved, err := a.Dao.GetBasketById(basketId)
	if err != nil {
		return errors.New("something went wrong"), http.StatusInternalServerError
	}
	if basketRetrieved == nil {
		return errors.New("basket not found"), http.StatusNotFound
	}

	product, err := a.Dao.GetProductById(productId)
	if err != nil {
		return errors.New("something went wrong"), http.StatusInternalServerError
	}
	if product == nil {
		return errors.New("product not found"), http.StatusNotFound
	}

	if product.Quantity < quantity {
		return errors.New("not enough products"), http.StatusBadRequest
	}

	productInBasket, err := a.Dao.GetProductFromBasketById(basketId, productId)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	if productInBasket == nil {
		err := a.createProductInBasket(basketId, productId, quantity)
		if err != nil {
			return errors.New("something went wrong"), http.StatusInternalServerError
		}
	} else {
		err = a.updateProductInBasket(*productInBasket, quantity)
		if err != nil {
			return err, http.StatusInternalServerError
		}
	}

	err = a.decreaseProductQuantity(product, quantity)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return nil, http.StatusOK
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
