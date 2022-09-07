package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"webserver/models"
)

func (c *Controller) GetProductById(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var idReqest models.IdRequest
	err = json.Unmarshal(body, &idReqest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	productRetrieved, err := c.Dao.GetProductById(idReqest.ID)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, productRetrieved, http.StatusOK)
}
