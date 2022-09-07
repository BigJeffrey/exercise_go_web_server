package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"webserver/models"
)

func (c *Controller) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var productRequest models.Product
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}
	err = c.Dao.UpdateProduct(productRequest)

	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, "product updated", http.StatusOK)
}
