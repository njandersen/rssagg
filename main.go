package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in environment variables")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr: ":" + port,
	}

	log.Printf("Server is running on port %v", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT is", port)
}