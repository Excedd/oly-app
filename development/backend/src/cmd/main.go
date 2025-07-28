package main

import (
	"context"
	"log"
	PersonalRecordHandler "oly-backend/application/handler/PersonalRecordHandler"
	TrainingHandler "oly-backend/application/handler/TrainingHandler"

	PersonalRecordRepository "oly-backend/repository"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
	mu         sync.Mutex
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = client.Disconnect(ctx)
	}()

	collection = client.Database("develop").Collection("prs")
	PersonalRecordRepository.InitPRRepository(collection)

	app := fiber.New()

	//app.Post("/api/prs", handleSetPRs)
	app.Get("/api/prs", PersonalRecordHandler.HandleGetPRs)

	//app.Post("/api/prs/calculateRound", handleSetPRs)

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello with MongoDB backend!")
	})

	app.Get("/training/calculate", TrainingHandler.HandlePlatesForRounds())

	log.Println("Backend listening on :8080")
	log.Fatal(app.Listen(":8080"))
}
