package services

import (
	"encoding/json"
	"net/http"
	"potato-module/contracts"
	"potato-module/models"
)

func (s *Services) CreateType(w http.ResponseWriter, r *http.Request) {
	var input contracts.CreateSizeBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := models.Type{
		Name: input.Name,
	}

	if err := s.DB.Create(&t).Error; err != nil {
		http.Error(w, "Error creating type: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func (s *Services) GetAllTypes(w http.ResponseWriter, r *http.Request) {

	var t []models.Type
	if err := s.DB.Find(&t).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}
