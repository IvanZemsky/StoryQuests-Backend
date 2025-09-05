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

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "stories-backend/docs"
)

// @title           Story Quests API
// @version         1.0
// @description     API for StoryQuests web site
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
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

	setupSwagger(router)
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

func setupSwagger(router *gin.Engine) {
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
