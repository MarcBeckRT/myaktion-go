package main

import (
	
	"os"
	
	log "github.com/sirupsen/logrus"
	
	"net/http"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/handler"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting My-Aktion API server")
	router := mux.NewRouter()
	router.HandleFunc("/health", handler.Health).Methods("GET")
	router.HandleFunc("/campaign", handler.CreateCampaign).Methods("POST")
	router.HandleFunc("/campaigns", handler.GetCampaigns).Methods("GET")
	router.HandleFunc("/donation", handler.AddDonation).Methods("POST")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
	
}

func init() {
	// init logger
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
	log.Info("Log level not specified, set default to: INFO")
	log.SetLevel(log.InfoLevel)
	return
	}
	log.SetLevel(level)
	}
