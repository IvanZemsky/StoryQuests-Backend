package main

import (
	"log"
	"stories-backend/config"
	"stories-backend/internal/handlers"
	"stories-backend/internal/repository"
	"stories-backend/internal/service"
	"stories-backend/pkg/db"

	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := config.ReadConfig("config/config.yml")
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	URI := "mongodb://" + config.Database.Host + ":" + fmt.Sprint(config.Database.Port)

	client, err := db.NewMongoDB(URI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	r := gin.Default()

	storyRepository := repository.NewStoryRepository(client.Database(config.Database.Name))
	storyService := service.NewStoryService(storyRepository)
	handlers.NewStoryHandler(r, storyService)

	r.Run(":" + fmt.Sprint(config.Port))

}
