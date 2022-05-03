package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/models"
	"github.com/MarcBeckRT/myaktion-go/src/myaktion/service"
)

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign models.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Printf("Can't serialize request body to campaign struct: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(&campaign); err != nil {
		log.Printf("Error callingserviceCreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaign); err != nil {
		log.Printf("Failureencodingvalueto JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Printf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaigns); err != nil {
		log.Printf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
