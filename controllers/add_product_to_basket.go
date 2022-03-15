package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"webserver/models"
)

func (c *Controller) AddProductToBasket(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var productRequest models.ProductToBasketRequest
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	err, code := c.AddProductToBasketService.Add(productRequest.BasketId, productRequest.ProductId, productRequest.Quantity)
	if err != nil {
		ApiError(w, err.Error(), code)

	}

	ApiSuccess(w, "product added to basket", http.StatusCreated)
}
