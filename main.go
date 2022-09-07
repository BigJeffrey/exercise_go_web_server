package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"webserver/controllers"
	"webserver/services"

	"github.com/gorilla/mux"

	postgresqldao "webserver/dao/postgresql"
)

func handlers(c *controllers.Controller) {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/product", c.GetAllProducts).Methods("GET")
	myRouter.HandleFunc("/product/id", c.GetProductById).Methods("GET")
	myRouter.HandleFunc("/product", c.CreateProduct).Methods("POST")
	myRouter.HandleFunc("/product", c.UpdateProduct).Methods("PUT")
	myRouter.HandleFunc("/product", c.DeleteProduct).Methods("DELETE")
	myRouter.HandleFunc("/product/filter", c.FilterProduct).Methods("GET")
	myRouter.HandleFunc("/productbyprice", c.GetProductByPrice).Methods("GET")
	myRouter.HandleFunc("/basket", c.CreateBasket).Methods("POST")
	myRouter.HandleFunc("/basket/product", c.AddProductToBasket).Methods("POST")
	myRouter.HandleFunc("/basket", c.GetAllBaskets).Methods("GET")
	myRouter.HandleFunc("/basket/product", c.DeleteProductInBasket).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	ctx := context.Background()

	conString := fmt.Sprintf("postgres://postgres:%s@host.docker.internal:5432/postgres?sslmode=disable", os.Getenv("POSTGRES_PASSWORD"))

	dao := postgresqldao.NewPostgreSql(ctx, conString)
	defer dao.Disconnect()

	addProductToBasketService := &services.AddProductToBasketService{
		Dao: dao,
	}

	controller := &controllers.Controller{Dao: dao, AddProductToBasketService: addProductToBasketService}

	handlers(controller)
}
