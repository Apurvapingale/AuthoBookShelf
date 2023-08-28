package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Apurvapingale/book-store/package/config"
	"github.com/Apurvapingale/book-store/package/models"
	"github.com/Apurvapingale/book-store/package/routes"
)

func main() {
	r := mux.NewRouter()
	db := config.ConnectDB()
	//creating table in database
	fmt.Println("Creating table")

	err := db.AutoMigrate(&models.User{}, &models.Book{}, &models.CartItem{}, &models.Order{}, &models.OrderDetail{}, &models.RatingReview{})
	if err != nil {
		fmt.Println("Error while creating table")
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Book Store")
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Welcome to Book Store"}`))

	})

	routes.RegisterBookStoreRoutes(r)
	routes.UserRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is listening on port 9010")
	if err := http.ListenAndServe("localhost:9010", r); err != nil {
		log.Fatal(err)
	}

}
