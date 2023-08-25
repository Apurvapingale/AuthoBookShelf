package routes

import (
	"github.com/Apurvapingale/book-store/package/controllers"
	"github.com/Apurvapingale/book-store/package/middleware"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	//create the middleware for admin
    // router.HandleFunc("/admin/login", controllers.UserLogin).Methods("POST")
	// router.HandleFunc("/admin/register", controllers.AdminRegister).Methods("POST")
    bookRoutes := router.PathPrefix("/book").Subrouter()

	bookRoutes.Use(middleware.ValidateAdmin)

	bookRoutes.HandleFunc("/", controllers.CreateBook).Methods("POST")

	bookRoutes.HandleFunc("/", controllers.GetBook).Methods("GET")

	bookRoutes.HandleFunc("/{bookId}", controllers.GetBookById).Methods("GET")

	// bookRoutes.HandleFunc("/{bookId}", controllers.UpdateBook).Methods("PUT")

	bookRoutes.HandleFunc("/{bookId}", controllers.DeleteBook).Methods("DELETE")
}






