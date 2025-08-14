package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Scene struct {
	ID          bson.ObjectID `bson:"_id" json:"id"`
	Number      int    `bson:"number" json:"number"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	StoryID     bson.ObjectID  `bson:"storyId" json:"storyId"`
	Img         string `bson:"img" json:"img"`
	// default | end
	Type    string        `bson:"type" json:"type"`
	Answers []SceneAnswer `bson:"answers" json:"answers"`
}

type SceneAnswer struct {
	ID              string `bson:"id" json:"id"`
	Text            string `bson:"text" json:"text"`
	NextSceneNumber int    `bson:"nextSceneNumber" json:"nextSceneNumber"`
}

type SceneRepository interface {
	FindByStoryID(storyID bson.ObjectID) ([]Scene, error)
}

type SceneService interface {
	FindByStoryID(id bson.ObjectID) ([]Scene, error)
}