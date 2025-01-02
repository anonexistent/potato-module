package services

import (
	"encoding/json"
	"net/http"
	"potato-module/contracts"
	"potato-module/models"
	"strings"

	"github.com/google/uuid"
)

func (s *Services) InitCart(w http.ResponseWriter, r *http.Request) {
	var input contracts.CreateCart
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart := models.Cart{
		Payload: input.Payload,
	}

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

	var input contracts.CreateCart
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var result string

	if strings.Contains(cart.Payload, ", "+input.Payload) {
		result = strings.ReplaceAll(cart.Payload, ", "+input.Payload, "")
	} else if strings.Contains(cart.Payload, ","+input.Payload) {
		result = strings.ReplaceAll(cart.Payload, ","+input.Payload, "")
	} else if strings.Contains(cart.Payload, input.Payload) {
		result = strings.ReplaceAll(cart.Payload, input.Payload, "")
	} else {
		http.Error(w, "product not found"+result, http.StatusBadRequest)
		return
	}

	cart.Payload = result
	if err := s.DB.Save(&cart).Error; err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

func (s *Services) PushCart(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var cart models.Cart
	if err := s.DB.Where("id = ?", id).First(&cart).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	var input contracts.CreateCart
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart.Payload += input.Payload

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
	if err := s.DB.First(&cart, uuid).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}
