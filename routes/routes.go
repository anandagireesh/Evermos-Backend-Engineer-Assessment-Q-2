package routes

import (
	"github.com/anandagireesh/Evermos-Backend-Engineer-Assessment-Q-2/controllers"
	"github.com/gorilla/mux"
)

func MainRoutes() *mux.Router {

	r := mux.NewRouter().StrictSlash(false)

	//Register products

	r.HandleFunc("/api/product/addproduct", controllers.AddProducts).Methods("POST")

	//buy products

	r.HandleFunc("/api/product/buyproduct", controllers.BuyProduct).Methods("POST")

	return r

}
