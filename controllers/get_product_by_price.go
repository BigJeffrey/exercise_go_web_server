package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webserver/models"
)

func (c *Controller) GetProductByPrice(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var priceSelectRequest models.PriceSelectRequest
	err = json.Unmarshal(body, &priceSelectRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	products, err := c.Dao.GetProductsByPrice(priceSelectRequest)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, products, http.StatusOK)
}
