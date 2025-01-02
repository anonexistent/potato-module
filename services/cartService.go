package services

import (
	"encoding/json"
	"net/http"
	"potato-module/contracts"
	"potato-module/models"

	"github.com/google/uuid"
)

func (s *Services) InitCart(w http.ResponseWriter, r *http.Request) {
	cart := models.Cart{}

	if err := s.DB.Create(&cart).Error; err != nil {
		http.Error(w, "Error creating cart: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (s *Services) RemoveFrom(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var cart models.Cart
	if err := s.DB.First(&cart, uuid).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var input contracts.PositionIdBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.DB.Where("id = ?", input.ID).Delete(&models.CartPosition{}).Error; err != nil {
		http.Error(w, "Error deleting position: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i, position := range cart.Positions {
		if position.ID == input.ID {
			cart.Positions = append(cart.Positions[:i], cart.Positions[i+1:]...)
			break
		}
	}

	if err := s.DB.Save(&cart).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (s *Services) PushCart(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cart models.Cart
	if err := s.DB.First(&cart, uuid).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var input contracts.CreateCart
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	input.Position.CartId = uuid

	if err := s.DB.Create(&input.Position).Error; err != nil {
		http.Error(w, "Error creating position: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.DB.Save(&cart).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (s *Services) GetCart(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var cart models.Cart
	if err := s.DB.Preload("Positions").First(&cart, uuid).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
