package models

import "github.com/google/uuid"

type SearcheRequest struct {
	Searche string
}

type PriceSelectRequest struct {
	PriceCondition string
	Price          float64
}

type ProductToBasketRequest struct {
	ProductId uuid.UUID
	Quantity  int
	BasketId  uuid.UUID
}

type IdRequest struct {
	ID uuid.UUID
}

type DeleteProductInBasketRequest struct {
	ProductID uuid.UUID
	BasketID  uuid.UUID
}
