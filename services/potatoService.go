package services

import (
	"encoding/json"
	_ "fmt"
	"net/http"
	"potato-module/contracts"
	"potato-module/models"
	"strconv"
)

// CreatePotato handles the creation of a new potato
func (s *Services) CreatePotato(w http.ResponseWriter, r *http.Request) {
	var input contracts.CreatePotatoBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var sizes []models.Size
	var types []models.Type
	var cats []models.Category

	if err := s.DB.Where("id IN ?", input.Sizes).Find(&sizes).Error; err != nil {
		http.Error(w, "Error finding sizes: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.DB.Where("id IN ?", input.Types).Find(&types).Error; err != nil {
		http.Error(w, "Error finding types: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := s.DB.Where("id IN ?", input.Categories).Find(&cats).Error; err != nil {
		http.Error(w, "Error finding types: "+err.Error(), http.StatusBadRequest)
		return
	}

	potato := models.Potato{
		Price:      input.Price,
		Title:      input.Title,
		Img:        input.Img,
		Rating:     input.Rate,
		Sizes:      sizes,
		Types:      types,
		Categories: cats,
	}

	if err := s.DB.Create(&potato).Error; err != nil {
		http.Error(w, "Error creating potato: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(potato)
}

// GetPotatoByID handles fetching a potato by its ID
func (s *Services) GetPotatoByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	pid, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	var potato models.Potato
	if err := s.DB.Preload("Types").Preload("Sizes").First(&potato, pid).Error; err != nil {
		http.Error(w, "Potato not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(potato)
}

// GetAllPotatoes handles fetching all potatoes
func (s *Services) GetAllPotatoes(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры пагинации из запроса
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")

	query := s.DB.Preload("Types").Preload("Sizes").Preload("Categories")
	sortField := r.URL.Query().Get("sort")
	categoryFilter := r.URL.Query().Get("category")
	// Если указан фильтр по категориям, добавляем его в запрос
	if categoryFilter != "" {
		query = query.Joins("JOIN potato_categoris ON potato_categoris.potato_id = potatos.id").
			Joins("JOIN categories ON categories.id = potato_categoris.category_id").
			Where("categories.id = ?", categoryFilter)
	}
	// Если указано поле для сортировки, добавляем его в запрос
	if sortField != "" {
		switch sortField {
		case "Title", "Rating", "Price":
			query = query.Order(sortField)
		default:
			http.Error(w, "Invalid sort field", http.StatusBadRequest)
			return
		}
	}

	// Устанавливаем значения по умолчанию, если параметры не указаны
	if page == "" {
		page = "1"
	}
	if pageSize == "" {
		pageSize = "6"
	}

	// Преобразуем параметры в целые числа
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		http.Error(w, "Invalid page number", http.StatusBadRequest)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt < 1 {
		http.Error(w, "Invalid page size", http.StatusBadRequest)
		return
	}

	// Вычисляем смещение
	offset := (pageInt - 1) * pageSizeInt

	// Подсчитываем общее количество записей
	var totalRecords int64
	if err := s.DB.Model(&models.Potato{}).Count(&totalRecords).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Вычисляем общее количество страниц
	totalPages := (totalRecords + int64(pageSizeInt) - 1) / int64(pageSizeInt)

	var potatoes []models.Potato

	if err := query.Limit(pageSizeInt).Offset(offset).Find(&potatoes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем структуру для ответа
	response := struct {
		Potatoes   []models.Potato `json:"potatoes"`
		TotalPages int64           `json:"totalPages"`
	}{
		Potatoes:   potatoes,
		TotalPages: totalPages,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
