package controllers

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"webserver/consts"
	"webserver/interfaces"
	"webserver/models"
	"webserver/services"
)

type ProductController struct {
	PG  interfaces.PostgresqlInterface
	Log *zap.SugaredLogger
	PS  *services.DeleteProductService
}

func NewProductController(PG interfaces.PostgresqlInterface, log *zap.SugaredLogger, PS *services.DeleteProductService) *ProductController {
	return &ProductController{PG: PG, Log: log, PS: PS}
}

func (p *ProductController) GetProductById(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	var idReqest models.IdRequest
	err = json.Unmarshal(body, &idReqest)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	productRetrieved, err := p.PG.GetProductById(idReqest.ID)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		p.Log.Error(err.Error())
		return
	}

	ReturnJSON(w, productRetrieved, http.StatusOK)
}

func (p *ProductController) GetProductByPrice(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	var priceSelectRequest models.PriceSelectRequest
	err = json.Unmarshal(body, &priceSelectRequest)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	products, err := p.PG.GetProductsByPrice(priceSelectRequest)
	if err != nil {
		ApiError(w, "something went wrong", http.StatusInternalServerError)
		p.Log.Error(err.Error())
		return
	}

	ReturnJSON(w, products, http.StatusOK)
}

func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	var newProduct models.Product
	err = json.Unmarshal(body, &newProduct)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	id, err := p.PG.CreateProduct(newProduct)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		p.Log.Error(err.Error())
		return
	}

	ApiSuccess(w, id.String(), http.StatusCreated)
}

func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	var productRequest models.Product
	err = json.Unmarshal(body, &productRequest)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	err = p.PG.UpdateProduct(productRequest)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		p.Log.Error(err.Error())
		return
	}

	ApiSuccess(w, "product updated", http.StatusOK)
}

func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	var idRequest models.IdRequest
	err = json.Unmarshal(body, &idRequest)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	err = p.PS.DeleteProductService(idRequest)
	if err != nil {
		if err.Error() == consts.ProductNotFoundError {
			ApiError(w, err.Error(), http.StatusNotFound)
			return
		}

		ApiError(w, fmt.Sprintf("something went wrong: %s", err.Error()), http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	ApiSuccess(w, fmt.Sprintf("product %v deleted", idRequest.ID), http.StatusOK)
}

func (p *ProductController) FilterProduct(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	var search models.SearcheRequest
	err = json.Unmarshal(body, &search)
	if err != nil {
		ApiError(w, consts.CantDecodeError, http.StatusBadRequest)
		p.Log.Error(err.Error())
		return
	}

	product, err := p.PG.FilterProduct(search.Searche)
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		p.Log.Error(err.Error())
		return
	}

	ReturnJSON(w, product, http.StatusOK)
}

func (p *ProductController) GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	products, err := p.PG.GetProducts()
	if err != nil {
		ApiError(w, err.Error(), http.StatusInternalServerError)
		p.Log.Error(err.Error())
		return
	}

	ReturnJSON(w, products, http.StatusOK)
}
