package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Story struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	AuthorID    bson.ObjectID `bson:"authorId" json:"authorId"`
	SceneCount  int           `bson:"sceneCount" json:"sceneCount"`
	Img         string        `bson:"img"  json:"img"`
	Likes       int           `bson:"likes"  json:"likes"`
	Date        time.Time     `bson:"date,omitempty" json:"date"`
	Passes      int           `bson:"passes" json:"passes"`
	Tags        []string      `bson:"tags" json:"tags"`
}

type StoryService interface {
	Find(filters StoryFilters) ([]Story, error)
	FindByID(id string) (Story, error)
}

type StoryRepository interface {
	Find(filters StoryFilters) ([]Story, error)
	FindByID(id bson.ObjectID) (Story, error)
}
