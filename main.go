package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"potato-module/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Size represents the size of potatoo
type Size struct {
	ID      uint             `json:"id" gorm:"primaryKey"`
	Name    string           `json:"name"`
	Potatos []*models.Potato `gorm:"many2many:potato_sizes;"`
}

// CreatePotatoBody represents the input structure for creating a potato
type CreatePotatoBody struct {
	Img   string `json:"img"`
	Price uint   `json:"price"`
	Title string `json:"title"`
	Types []uint `json:"types"`
	Sizes []uint `json:"sizes"`
}

var db *gorm.DB

func initializeData() {
	// Инициализация типов картошки
	types := []models.Type{
		{Name: "молодая"},
		{Name: "старая"},
	}

	// Инициализация размеров картошки
	sizes := []Size{
		{Name: "бэби"},
		{Name: "медиум"},
		{Name: "босс"},
	}

	// Сохранение данных в базе
	for _, t := range types {
		db.FirstOrCreate(&t)
	}
	for _, s := range sizes {
		db.FirstOrCreate(&s)
	}
}

func getDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
}

func main() {
	if os.Getenv("MODE") != "prod" {
		// Загружаем переменные окружения из файла Development.env
		//
		err := godotenv.Load("Development.env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// Подключаемся к PostgreSQL без указания базы данных
	connStr := "host=localhost user=postgres password=sa port=5432 sslmode=disable"
	temp_db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}

	// Создаем базу данных, если она не существует
	_, err = temp_db.Exec("CREATE DATABASE \"potato-module-db\"")
	if err != nil {
		fmt.Printf("Error creating database: %v\n", err)
	} else {
		fmt.Println("Database created successfully")
	}
	temp_db.Close()

	dsn := getDSN()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Potato{}, &models.Type{}, &models.Size{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	initializeData()

	// Initialize the router
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)

	// Define routes
	r.Post("/potatoes/create", createPotato)
	r.Get("/potatoes/{id}", getPotatoByID)
	r.Get("/potatoes/list", getAllPotatoes)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// createPotato handles the creation of a new potato
func createPotato(w http.ResponseWriter, r *http.Request) {
	var input CreatePotatoBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var potato = models.Potato{
		Price: input.Price,
		Title: input.Title,
		Img:   input.Img,
	}

	result := db.Create(&potato)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	var s []models.Size
	var t []models.Type
	db.Where("id in ?", input.Sizes).Find(&s)
	db.Where("id in ?", input.Types).Find(&t)

	potato.Sizes = s
	potato.Types = t

	db.Save(&potato)

	json.NewEncoder(w).Encode(potato)
}

// getPotatoByID handles fetching a potato by its ID
func getPotatoByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var potato models.Potato
	if err := db.Preload("Types").Preload("Sizes").First(&potato, id).Error; err != nil {
		http.Error(w, "Potato not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(potato)
}

// getAllPotatoes handles fetching all potatoes
func getAllPotatoes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var potatoes []models.Potato
	if err := db.Preload("Types").Preload("Sizes").Find(&potatoes).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(r.RemoteAddr)

	json.NewEncoder(w).Encode(potatoes)
}
