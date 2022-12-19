package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"webserver/controllers"
	"webserver/persistence"
	"webserver/router"
	"webserver/services"
)

func main() {
	ctx := context.Background()

	conString := fmt.Sprintf("postgres://postgres:%s@host.docker.internal:5432/postgres?sslmode=disable", os.Getenv("POSTGRES_PASSWORD"))

	postgresqlService := persistence.NewPostgreSql(ctx, conString)
	defer postgresqlService.Disconnect()

	addProductToBasketService := &services.AddProductToBasketService{
		Dao: postgresqlService,
	}

	basketController := controllers.NewBasketController(postgresqlService, addProductToBasketService)
	productController := controllers.NewProductController(postgresqlService)

	myRouter := router.Handlers(basketController, productController)

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
