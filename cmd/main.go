package main

import (
	"log"
	"net/http"
	"time"

	"github.com/qooodr/CRUD-APP/internal/repository/psql"
	"github.com/qooodr/CRUD-APP/internal/service"
	"github.com/qooodr/CRUD-APP/internal/transport/rest"
	"github.com/qooodr/CRUD-APP/pkg/database"

	_ "github.com/lib/pq"
)

func main() {
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	booksRepo := psql.NewBooks(db)
	booksService := service.NewBooks(booksRepo)
	handler := rest.NewHandler(booksService)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRouter(),
	}

	log.Printf("SERVER STARTED AT", time.Now().Format(time.RFC3339))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
