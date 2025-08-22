package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type StoryLike struct {
	ID      bson.ObjectID `bson:"_id,omitempty" json:"id"`
	StoryID bson.ObjectID `bson:"storyId" json:"storyId"`
	UserID  bson.ObjectID `bson:"userId" json:"userId"`
}

type StoryLikeRepository interface {
	FindLike(storyID bson.ObjectID, userID bson.ObjectID) ([]LikeStoryResponse, error)
	AddLike(storyID bson.ObjectID, userID bson.ObjectID) error
	RemoveLike(storyID bson.ObjectID, userID bson.ObjectID) error
}
