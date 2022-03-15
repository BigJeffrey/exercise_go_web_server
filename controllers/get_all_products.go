package controllers

import (
	"net/http"
)

func (c *Controller) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.Dao.GetProducts()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, products, http.StatusOK)
}
