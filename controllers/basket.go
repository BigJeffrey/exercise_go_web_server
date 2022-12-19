package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"webserver/interfaces"
	"webserver/models"
	"webserver/services"
)

type BasketController struct {
	PG                        interfaces.PostgresqlInterface
	AddProductToBasketService *services.AddProductToBasketService
}

func NewBasketController(PG interfaces.PostgresqlInterface, addProductToBasketService *services.AddProductToBasketService) *BasketController {
	return &BasketController{PG: PG, AddProductToBasketService: addProductToBasketService}
}

func (b *BasketController) AddProductToBasket(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var productRequest models.ProductToBasketRequest
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	err, code := b.AddProductToBasketService.Add(productRequest.BasketId, productRequest.ProductId, productRequest.Quantity)
	if err != nil {
		ApiError(w, err.Error(), code)
		return
	}

	ApiSuccess(w, "product added to basket", http.StatusCreated)
}

func (b *BasketController) CreateBasket(w http.ResponseWriter, _ *http.Request) {
	id, err := b.PG.CreateBasket()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, id.String(), http.StatusCreated)
}

func (b *BasketController) DeleteProductInBasket(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	var productRequest models.DeleteProductInBasketRequest
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		return
	}

	productInBasket, err := b.PG.GetProductFromBasketById(productRequest.ProductID, productRequest.BasketID)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if productInBasket == nil {
		ApiError(w, "product is not in basket", http.StatusNotFound)
		return
	}

	_, err = b.PG.DeleteProductFromBasket(models.BasketProducts{
		BasketID:  productRequest.BasketID,
		ProductID: productRequest.ProductID,
	})
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, "deleted", http.StatusOK)
}

func (b *BasketController) GetAllBaskets(w http.ResponseWriter, _ *http.Request) {
	var basketSummary models.BasketSummary
	var basketsSummaryTab []models.BasketSummary

	baskets, err := b.PG.GetBaskets()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products, err := b.PG.GetProducts()
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
