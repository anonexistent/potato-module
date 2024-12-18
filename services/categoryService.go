package services

import (
	"encoding/json"
	"net/http"
	"potato-module/contracts"
	"potato-module/models"
)

func (s *Services) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var input contracts.CreateSizeBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c := models.Category{
		Name: input.Name,
	}

	if err := s.DB.Create(&c).Error; err != nil {
		http.Error(w, "Error creating size: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (s *Services) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	var ss []models.Category
	if err := s.DB.Find(&ss).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ss)
}
