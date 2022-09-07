package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"webserver/models"
)

func (c *Controller) DeleteProductInBasket(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var productRequest models.DeleteProductInBasketRequest
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	productInBasket, err := c.Dao.GetProductFromBasketById(productRequest.ProductID, productRequest.BasketID)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if productInBasket == nil {
		ApiError(w, "product is not in basket", http.StatusNotFound)
		return
	}

	_, err = c.Dao.DeleteProductFromBasket(models.BasketProducts{
		BasketID:  productRequest.BasketID,
		ProductID: productRequest.ProductID,
	})
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, "deleted", http.StatusOK)
}
