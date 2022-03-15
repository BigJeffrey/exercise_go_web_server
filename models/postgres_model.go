package models

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	Price       float64
	Quantity    int
	Status      string
}

type Basket struct {
	ID uuid.UUID
}

type BasketProducts struct {
	ID        uuid.UUID
	BasketID  uuid.UUID
	ProductID uuid.UUID
	Quantity  int
}
