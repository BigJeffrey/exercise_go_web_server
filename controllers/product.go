package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"webserver/interfaces"
	"webserver/models"
)

type ProductController struct {
	PG interfaces.PostgresqlInterface
}

func NewProductController(PG interfaces.PostgresqlInterface) *ProductController {
	return &ProductController{PG: PG}
}

func (p *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
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

	productRetrieved, err := p.PG.GetProductById(idReqest.ID)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, productRetrieved, http.StatusOK)
}

func (p *ProductController) GetProductByPrice(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	var priceSelectRequest models.PriceSelectRequest
	err = json.Unmarshal(body, &priceSelectRequest)
	if err != nil {
		ApiError(w, "can not decode request", http.StatusBadRequest)
		log.Println(err)
		return
	}

	products, err := p.PG.GetProductsByPrice(priceSelectRequest)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, products, http.StatusOK)
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
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

	id, err := p.PG.CreateProduct(newProduct)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, id.String(), http.StatusCreated)
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
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
	err = p.PG.UpdateProduct(productRequest)

	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, "product updated", http.StatusOK)
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
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

	productRetrieved, err := p.PG.GetProductById(idRequest.ID)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	if productRetrieved == nil {
		ApiError(w, "product not exists", http.StatusNotFound)
		return
	}

	_, err = p.PG.DeleteProduct(idRequest.ID)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	ApiSuccess(w, "deleted", http.StatusOK)
}

func (p *ProductController) FilterProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
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

	product, err := p.PG.FilterProduct(search.Searche)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, product, http.StatusOK)
}

func (p *ProductController) GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := p.PG.GetProducts()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ReturnJSON(w, products, http.StatusOK)
}
