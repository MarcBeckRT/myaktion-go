package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/service"
	"github.com/MarcBeckRT/myaktion-go/src/myaktion/models"
)

func AddDonation(w http.ResponseWriter, r *http.Request) {
	donation, err := getDonation(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.AddDonation(donation); err != nil {
		log.Printf("Error calling service AddDonation: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, donation)
}

func getDonation(r *http.Request) (*models.Donation, error) {
	var donation models.Donation
	err := json.NewDecoder(r.Body).Decode(&donation)
	if err != nil {
	log.Errorf("Can't serialize request body to donation struct: %v", err)
	return nil, err
	}
	return &donation, nil
}