package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webserver/models"
)

func (c *Controller) GetProductById(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var idReqest models.IdReqest
	err = json.Unmarshal(body, &idReqest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	productRetrived, err := c.Dao.GetProductById(idReqest.ID)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, productRetrived, http.StatusOK)
}
