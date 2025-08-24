package main

import (
	"context"
	"log"
	"net/http"
	"stories-backend/config"
	"stories-backend/internal/compose"
	"stories-backend/pkg/db/mongo"
	"time"

	"fmt"

	commonHandlers "stories-backend/internal/handlers/common"

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

	setupCORS(router, config)

	compose.InitModules(compose.InitModuleOptions{Client: client, Config: config, Router: router})

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

func setupCORS(router *gin.Engine, config *config.Config) {
	router.Use(commonHandlers.CORSMiddleware(config.Origin))

	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.Origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, Cookie, Set-Cookie")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Expose-Headers", "Set-Cookie")
		c.Status(http.StatusOK)
	})
}
