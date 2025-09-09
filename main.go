package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"potato-module/models"
	"potato-module/services"

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

var db *gorm.DB

// old
func initializeData(database *gorm.DB) {
	// Инициализация типов картошки
	types := []models.Type{
		{ID: 1, Name: "молодая"},
		{ID: 2, Name: "старая"},
	}

	// Инициализация размеров картошки
	sizes := []Size{
		{ID: 1, Name: "бэби"},
		{ID: 2, Name: "медиум"},
		{ID: 3, Name: "босс"},
	}

	// Сохранение данных в базе
	for _, t := range types {
		database.FirstOrCreate(&t)
	}
	for _, s := range sizes {
		database.FirstOrCreate(&s)
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

	ensureCreated()

	dsn := getDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Автоматически мигрировать схему
	if err := db.AutoMigrate(&models.Potato{}, &models.Type{}, &models.Size{}, &models.Category{}, &models.Cart{}, &models.CartPosition{}); err != nil {
		panic(fmt.Sprintln("Error during migration: %v\n", err))
	}

	// for release
	// // Получаем список всех таблиц
	// tables, err := db.Migrator().GetTables()
	// if err != nil {
	// 	log.Fatalf("failed to get tables: %v", err)
	// }
	// // Удаляем данные из каждой таблицы
	// for _, table := range tables {
	// 	if err := db.Exec("DELETE FROM " + table).Error; err != nil {
	// 		log.Printf("failed to clear table %s: %v", table, err)
	// 	} else {
	// 		log.Printf("cleared table %s", table)
	// 	}
	// }
	initializeData(db)

	// Initialize the router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,

		MaxAge: 300, // Maximum value not ignored by any of major browsers
	}))

	// Обработчик для OPTIONS-запросов
	r.Options("/*", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	var ss = &services.Services{DB: db}

	// Define routes
	r.Post("/potatoes/create", ss.CreatePotato)
	r.Get("/potatoes", ss.GetPotatoByID)
	r.Get("/potatoes/list", ss.GetAllPotatoes)

	r.Post("/sizes/create", ss.CreateSize)
	r.Get("/sizes/list", ss.GetAllSizes)

	r.Post("/types/create", ss.CreateType)
	r.Get("/types/list", ss.GetAllTypes)

	r.Post("/categories/create", ss.CreateCategory)
	r.Get("/categories/list", ss.GetAllCategories)

	r.Post("/cart/init", ss.InitCart)
	r.Get("/cart/get", ss.GetCart)
	r.Post("/cart/push", ss.PushCart)
	r.Delete("/cart/removeFrom", ss.RemoveFrom)

	// Start the server
	log.Println("Starting server on :54870")
	if err := http.ListenAndServe(":54870", r); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func ensureCreated() {
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
}
