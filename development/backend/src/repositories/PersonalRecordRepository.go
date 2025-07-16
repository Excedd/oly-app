package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	d "oly-backend/domain"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PRRepository interface {
	SavePRs(ctx context.Context, prs d.PersonalRecord) error
	GetPRs(ctx context.Context) (d.PersonalRecord, error)
}

var (
	prRepo PRRepository
)

func main() {
	// Conexión Mongo
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("prdb").Collection("prs")
	prRepo = repository.NewMongoPRRepository(collection)

	http.HandleFunc("/api/prs", prsHandler)
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello with Repository Pattern!")
	})

	fmt.Println("Backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func prsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handleSetPRs(w, r)
	case http.MethodGet:
		handleGetPRs(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleSetPRs(w http.ResponseWriter, r *http.Request) {
	var newPRs d.PersonalRecord

	if err := json.NewDecoder(r.Body).Decode(&newPRs); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := prRepo.SavePRs(ctx, newPRs)
	if err != nil {
		http.Error(w, "Failed to save PRs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPRs)
}

func handleGetPRs(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	prs, err := prRepo.GetPRs(ctx)
	if err != nil {
		http.Error(w, "No PRs found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prs)
}
