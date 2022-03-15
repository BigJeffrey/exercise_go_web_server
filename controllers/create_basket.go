package controllers

import (
	"net/http"
)

func (c *Controller) CreateBasket(w http.ResponseWriter, r *http.Request) {
	id, err := c.Dao.CreateBasket()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, id.String(), http.StatusCreated)
}
