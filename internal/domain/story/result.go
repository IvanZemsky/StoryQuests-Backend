package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type StoryResult struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   bson.ObjectID `json:"userId" bson:"userId"`
	StoryID  bson.ObjectID `json:"storyId" bson:"storyId"`
	SceneID  bson.ObjectID `json:"sceneId" bson:"sceneId"`
	Datetime bson.DateTime `json:"datetime,omitempty" bson:"datetime"`
}

type SetResultDTO struct {
	UserID   bson.ObjectID `json:"userId" bson:"userId"`
	StoryID  bson.ObjectID `json:"storyId" bson:"storyId"`
	SceneID  bson.ObjectID `json:"sceneId" bson:"sceneId"`
}

// for swaggo docs generation (error with bson.DateTime)
type GetStoryResult struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   bson.ObjectID `json:"userId" bson:"userId"`
	StoryID  bson.ObjectID `json:"storyId" bson:"storyId"`
	SceneID  bson.ObjectID `json:"sceneId" bson:"sceneId"`
	Datetime string `json:"datetime,omitempty" bson:"datetime"`
}
