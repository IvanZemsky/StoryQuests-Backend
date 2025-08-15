package compose

import (
	"stories-backend/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitModules(client *mongo.Client, config *config.Config, router *gin.Engine) {
	storyModule := InitStoryModule(client, config, router)
	InitSceneModule(client, config, router, storyModule.Repository)

}