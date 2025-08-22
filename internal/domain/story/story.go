package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Story struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	AuthorID    bson.ObjectID `bson:"author" json:"author"`
	SceneCount  int           `bson:"sceneCount" json:"sceneCount"`
	Img         string        `bson:"img"  json:"img"`
	Likes       int           `bson:"likes"  json:"likes"`
	Date        time.Time     `bson:"date,omitempty" json:"date"`
	Passes      int           `bson:"passes" json:"passes"`
	Tags        []string      `bson:"tags" json:"tags"`
}

type StoryResponse struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Author      struct {
		ID    bson.ObjectID `bson:"_id" json:"id"`
		Login string        `bson:"login" json:"login"`
	} `bson:"author" json:"author"`
	SceneCount int       `bson:"sceneCount" json:"sceneCount"`
	Img        string    `bson:"img"  json:"img"`
	Likes      int       `bson:"likes"  json:"likes"`
	Date       time.Time `bson:"date,omitempty" json:"date"`
	Passes     int       `bson:"passes" json:"passes"`
	Tags       []string  `bson:"tags" json:"tags"`
	IsLiked    bool      `bson:"isLiked" json:"isLiked"`
}

// type StoryResponse struct {
// 	Story  `json:",inline"`
// 	Author struct {
// 		ID    string `bson:"_id" json:"id"`
// 		Login string `bson:"login" json:"login"`
// 	} `bson:"author" json:"author"`
// }

type StoryService interface {
	Find(filters StoryFilters) ([]StoryResponse, error)
	FindByID(id bson.ObjectID) (StoryResponse, error)
	Like(LikeStoryDTO) (LikeStoryResponse, error)
}

type StoryRepository interface {
	Find(filters StoryFilters) ([]StoryResponse, error)
	FindByID(id bson.ObjectID) (StoryResponse, error)
	StoryExists(id bson.ObjectID) (bool, error)
	Like(LikeStoryDTO) (LikeStoryResponse, error)
}
