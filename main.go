package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/msartioli/helpdesk-api-go/controllers"
	"github.com/msartioli/helpdesk-api-go/database"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		log.Fatal("variável DATABASE_URL não definida")
	}

	db, err := database.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	controllers.Health()
	controllers.Tickets(db)

	log.Println("API rodando em http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
