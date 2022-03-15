package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"webserver/models"
)

func (c *Controller) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var newProduct models.Product
	err = json.Unmarshal(body, &newProduct)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	id, err := c.Dao.CreateProduct(newProduct)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, id.String(), http.StatusCreated)
}
