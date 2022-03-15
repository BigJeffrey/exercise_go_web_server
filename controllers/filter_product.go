package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"webserver/models"
)

func (c *Controller) FilterProduct(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var search models.SearcheRequest
	err = json.Unmarshal(body, &search)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	product, err := c.Dao.FilterProduct(search.Searche)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, product, http.StatusOK)
}
