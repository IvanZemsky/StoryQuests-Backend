package main

import (
	"context"
	"log"
	"stories-backend/config"
	storyDomain "stories-backend/internal/domain/story"
	"stories-backend/internal/handlers"
	"stories-backend/internal/repository"
	"stories-backend/internal/service"
	"stories-backend/pkg/db/mongo"
	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	config := readConfig("config/config.yml")

	URI := db.GetConnectionString(config)
	client := connectDB(URI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("An error occurred while disconnecting from MongoDB: %v", err)
		}
	}()

	router := gin.Default()

	storyRepo := initStoryModule(client, config, router)
	initSceneModule(client, config, router, *storyRepo)

	router.Run(":" + fmt.Sprint(config.Port))
}

func readConfig(path string) *config.Config {
	config, err := config.ReadConfig(path)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}
	return config
}

func connectDB(URI string) *mongo.Client {
	client, err := db.NewMongoDB(URI)
	if err != nil {
		log.Fatalf("Failed to connect to data base: %v", err)
	}
	return client
}

func initStoryModule(client *mongo.Client, config *config.Config, router *gin.Engine) *storyDomain.StoryRepository {
	storyRepository := repository.NewStoryRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("stories"),
	)
	storyService := service.NewStoryService(storyRepository)
	handlers.NewStoryHandler(router, storyService)

	return &storyRepository
}

func initSceneModule(client *mongo.Client,
	config *config.Config,
	router *gin.Engine,
	storyRepo storyDomain.StoryRepository,
) {
	sceneRepository := repository.NewSceneRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("scenes"),
	)
	sceneService := service.NewSceneService(sceneRepository, storyRepo)
	handlers.NewSceneHandler(router, sceneService)
}
