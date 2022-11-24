package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	port := ":7000"

	mux := http.NewServeMux()
	//define routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	fmt.Println("Started server on 7000 port")
	//starting server
	log.Fatal(http.ListenAndServe(port, mux))
}
