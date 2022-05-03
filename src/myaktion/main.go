package main

import (
	"log"
	"net/http"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/handler"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methodes("GET")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
