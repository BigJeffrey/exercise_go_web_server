package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"webserver/controllers"
	"webserver/persistence"
	"webserver/router"
	"webserver/services"
)

var (
	logger            *zap.SugaredLogger
	myRouter          *mux.Router
	postgresqlService *persistence.Postgresql
)

func init() {
	// logger
	zapLogger, _ := zap.NewDevelopment()
	defer zapLogger.Sync()
	logger = zapLogger.Sugar()
	logger.Info("sugar logger initialized")

	// database
	ctx := context.Background()
	conString := fmt.Sprintf("postgres://postgres:%s@host.docker.internal:5432/postgres?sslmode=disable", os.Getenv("POSTGRES_PASSWORD"))
	postgresqlService = persistence.NewPostgresql(ctx, conString)
	logger.Info("database initialized")
	//services
	addProductToBasketService := &services.AddProductToBasketService{Dao: postgresqlService}
	deleteProductService := &services.DeleteProductService{Log: logger, PG: postgresqlService}

	//controllers
	basketController := controllers.NewBasketController(postgresqlService, addProductToBasketService, logger)
	productController := controllers.NewProductController(postgresqlService, logger, deleteProductService)

	myRouter = router.Handlers(basketController, productController)
}

func main() {
	defer postgresqlService.Disconnect()

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
