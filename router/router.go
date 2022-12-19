package router

import (
	"github.com/gorilla/mux"
	"webserver/controllers"
)

func Handlers(b *controllers.BasketController, p *controllers.ProductController) *mux.Router {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/product", p.GetAllProducts).Methods("GET")
	myRouter.HandleFunc("/product/id", p.GetProductById).Methods("GET")
	myRouter.HandleFunc("/product", p.CreateProduct).Methods("POST")
	myRouter.HandleFunc("/product", p.UpdateProduct).Methods("PUT")
	myRouter.HandleFunc("/product", p.DeleteProduct).Methods("DELETE")
	myRouter.HandleFunc("/product/filter", p.FilterProduct).Methods("GET")
	myRouter.HandleFunc("/productbyprice", p.GetProductByPrice).Methods("GET")
	myRouter.HandleFunc("/basket", b.CreateBasket).Methods("POST")
	myRouter.HandleFunc("/basket/product", b.AddProductToBasket).Methods("POST")
	myRouter.HandleFunc("/basket", b.GetAllBaskets).Methods("GET")
	myRouter.HandleFunc("/basket/product", b.DeleteProductInBasket).Methods("DELETE")

	return myRouter
}
