package routes

import (
	"github.com/Apurvapingale/book-store/package/controllers"
	"github.com/Apurvapingale/book-store/package/middleware"
	"github.com/gorilla/mux"
)

var UserRoutes = func(router *mux.Router) {
	//create the middleware for admin
	userRoutes := router.PathPrefix("/user").Subrouter()

	router.HandleFunc("/user/register", controllers.RegisterUser).Methods("POST")

	router.HandleFunc("/user/login", controllers.LoginUser).Methods("POST")
	userRoutes.Use(middleware.ValidateUser)

	userRoutes.HandleFunc("/{userId}", controllers.GetUserData).Methods("GET") //get single user data by id

	userRoutes.HandleFunc("/{userId}", controllers.InActive).Methods("PUT") //update user status as inactive user

	userRoutes.HandleFunc("/{userId}", controllers.DeleteUser).Methods("DELETE") //update user status to deleted
	userRoutes.HandleFunc("/addreview", controllers.AddReview).Methods("POST")

	adminRoutes := router.PathPrefix("/admin").Subrouter()
	adminRoutes.Use(middleware.ValidateAdmin)

	adminRoutes.HandleFunc("/admin/register", controllers.RegisterAdmin).Methods("POST")

	cartRoutes := router.PathPrefix("/cart").Subrouter()
	cartRoutes.Use(middleware.ValidateUser)

	cartRoutes.HandleFunc("/addtoorder", controllers.AddToOrder).Methods("POST")
	cartRoutes.HandleFunc("/addtocart", controllers.AddToCart).Methods("POST")
	cartRoutes.HandleFunc("/decrement", controllers.DecrementCart).Methods("POST")
	cartRoutes.HandleFunc("/removeall", controllers.RemoveAllItemsFromCart).Methods("GET")
	cartRoutes.HandleFunc("/myorders", controllers.GetMyOrders).Methods("GET")

}
