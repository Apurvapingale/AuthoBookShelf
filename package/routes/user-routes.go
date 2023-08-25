package routes

import (
	"github.com/Apurvapingale/book-store/package/controllers"
	"github.com/Apurvapingale/book-store/package/middleware"

	"github.com/gorilla/mux"
)

var UserRoutes = func(router *mux.Router) {
	//create the middleware for admin
    bookRoutes := router.PathPrefix("/user").Subrouter()

	bookRoutes.HandleFunc("/login", controllers.LoginUser).Methods("POST")
	bookRoutes.Use(middleware.ValidateUser)
	
}