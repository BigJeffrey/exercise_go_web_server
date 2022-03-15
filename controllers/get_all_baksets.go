package controllers

import (
	"net/http"
	"webserver/models"

	"github.com/google/uuid"
)

func (c *Controller) GetAllBaskets(w http.ResponseWriter, r *http.Request) {
	var basketSummary models.BasketSummary
	var basketsSummaryTab []models.BasketSummary

	baskets, err := c.Dao.GetBaskets()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := c.Dao.GetProducts()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var mapProductCount = make(map[uuid.UUID]int)
	var mapTotalPrice = make(map[uuid.UUID]float64)

	for _, basket := range baskets {
		for _, product := range products {
			if basket.ProductID == product.ID {
				mapProductCount[basket.BasketID] += basket.Quantity
				mapTotalPrice[basket.BasketID] += float64(basket.Quantity) * product.Price
			}
		}
	}

	for key, val := range mapProductCount {
		basketSummary = models.BasketSummary{
			BasketId:      key,
			CountProducts: val,
			TotalPrice:    mapTotalPrice[key],
		}
		basketsSummaryTab = append(basketsSummaryTab, basketSummary)
	}

	ReturnJSON(w, basketsSummaryTab, http.StatusOK)
}
