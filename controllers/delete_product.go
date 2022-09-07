package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"webserver/models"
)

func (c *Controller) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var idRequest models.IdRequest
	err = json.Unmarshal(body, &idRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	productRetrieved, err := c.Dao.GetProductById(idRequest.ID)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	if productRetrieved == nil {
		ApiError(w, "product not exists", http.StatusNotFound)
		return
	}

	_, err = c.Dao.DeleteProduct(idRequest.ID)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, "deleted", http.StatusOK)
}
