package services

import (
	"encoding/json"
	"net/http"
	"potato-module/models"
)

func (s *Services) GetAllSizes(w http.ResponseWriter, r *http.Request) {

	var ss []models.Size
	if err := s.DB.Find(&ss).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ss)
}
