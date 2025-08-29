package domain

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateStoryDTO struct {
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	AuthorID    bson.ObjectID `bson:"author" json:"author"`
	SceneCount  int           `bson:"sceneCount" json:"sceneCount"`
	Img         string        `bson:"img"  json:"img"`
	Tags        []string      `bson:"tags" json:"tags"`
}

type CreateStoryInfoBody struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AuthorID    string   `json:"author"`
	SceneCount  int      `json:"sceneCount"`
	Img         string   `json:"img"`
	Tags        []string `json:"tags"`
}
