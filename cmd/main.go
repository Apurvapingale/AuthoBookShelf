package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Apurvapingale/book-store/package/routes"
)
func main() {
	r := mux.NewRouter()
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
