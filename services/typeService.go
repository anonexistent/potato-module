package services

import (
	"encoding/json"
	"net/http"
	"potato-module/models"
)

func (s *Services) GetAllTypes(w http.ResponseWriter, r *http.Request) {

	var t []models.Type
	if err := s.DB.Find(&t).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}
