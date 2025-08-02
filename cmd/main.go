package main

import (
	"log"
	"stories-backend/internal/handlers"
	"stories-backend/internal/repository"
	"stories-backend/internal/service"
	"stories-backend/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	client, err := db.NewMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	storyRepository := repository.NewStoryRepository(client)
	storyService := service.NewStoryService(storyRepository)
	handlers.NewStoryHandler(r, storyService)

	r.Run(":8080")
	
}
