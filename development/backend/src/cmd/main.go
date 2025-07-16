package main

import (
	"context"
	"log"
	d "oly-backend/domain"
	"os"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo client (global para demo)
var (
	client     *mongo.Client
	collection *mongo.Collection
	mu         sync.Mutex
)

func main() {
	// Setup Mongo
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

	app := fiber.New()

	app.Post("/api/prs", handleSetPRs)
	app.Get("/api/prs", handleGetPRs)

	app.Post("/api/prs/calculateRound", handleSetPRs)

	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello with MongoDB backend!")
	})

	log.Println("Backend listening on :8080")
	log.Fatal(app.Listen(":8080"))
}

func handleSetPRs(c *fiber.Ctx) error {
	var newPRs d.PersonalRecord
	if err := c.BodyParser(&newPRs); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	mu.Lock()
	defer mu.Unlock()

	_, err := collection.DeleteMany(ctx, map[string]interface{}{})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to clear old PRs")
	}

	_, err = collection.InsertOne(ctx, newPRs)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to insert PRs")
	}

	return c.Status(fiber.StatusOK).JSON(newPRs)
}

func handleGetPRs(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result d.PersonalRecord
	err := collection.FindOne(ctx, map[string]interface{}{}).Decode(&result)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "No PRs found")
	}

	return c.Status(fiber.StatusOK).JSON(result)
}
