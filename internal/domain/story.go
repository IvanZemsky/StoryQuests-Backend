package domain

import (
	_ "time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Story struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"` 
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	// Author      primitive.ObjectID `bson:"author"`
	// SceneCount  int                `bson:"sceneCount"`
	// Img         string             `bson:"img"`
	// Likes       int                `bson:"likes"`
	// Date        time.Time          `bson:"date,omitempty"`
	// Passes      int                `bson:"passes"`
	// Tags        []string           `bson:"tags"`
}

type StoryService interface {
	Find() ([]Story, error)
}

type StoryRepository interface {
	Find() ([]Story, error)
}
