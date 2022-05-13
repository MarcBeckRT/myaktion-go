package handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/MarcBeckRT/myaktion-go/src/myaktion/models"
	"github.com/MarcBeckRT/myaktion-go/src/myaktion/service"

	"github.com/gorilla/mux"
)

type result struct {
	Success string `json:"success"`
}

func sendJson(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		log.Errorf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getId(r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 0)
	if err != nil {
		log.Errorf("Can't get ID from request: %v", err)
		return 0, err
	}
	return uint(id), nil
}