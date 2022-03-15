package main

import (
	"context"
	"log"
	"net/http"
	"webserver/controllers"

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

	conString := "postgres://postgres:mysecretpassword@host.docker.internal:5432/postgres?sslmode=disable"

	dao := postgresqldao.NewPostgreSql(ctx, conString)
	defer dao.Disconnect()

	controllers := &controllers.Controller{Dao: dao}

	handlers(controllers)
}
