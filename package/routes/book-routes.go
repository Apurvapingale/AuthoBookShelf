package routes

import (
	"github.com/Apurvapingale/book-store/package/controllers"
	"github.com/Apurvapingale/book-store/package/middleware"

	"github.com/gorilla/mux"
)

// registerbookstore function is used to register the routes
var RegisterBookStoreRoutes = func(router *mux.Router) {

	//all book routes have the group prefix /book
	bookRoutes := router.PathPrefix("/book").Subrouter()

	bookRoutes.Use(middleware.ValidateAdmin)

	//create the routes for book
	bookRoutes.HandleFunc("/", controllers.CreateBook).Methods("POST")

	bookRoutes.HandleFunc("/", controllers.GetBook).Methods("GET")

	bookRoutes.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")

	bookRoutes.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")

	bookRoutes.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
