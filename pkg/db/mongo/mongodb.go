package db

import (
	"context"
	"fmt"
	"log"
	"stories-backend/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func NewMongoDB(URI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

func GetConnectionString(config *config.Config) string {
	switch config.DBType {
	case "cluster":
		return fmt.Sprintf(
			"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority&appName=%s",
			config.Database.UserName,
			config.Database.Password,
			config.Database.ClusterCode,
			config.Database.Name,
			config.Database.ClusterName,
		)
	case "local":
		return fmt.Sprintf("mongodb://%s:%d", config.Database.Host, config.Database.Port)
	default:
		log.Fatal("Unknown connection type")
		return ""
	}
}
