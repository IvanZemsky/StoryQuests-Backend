package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	sceneDomain "stories-backend/internal/domain/scene"
)

type StoryService interface {
	Find(filters StoryFilters) ([]StoryResponse, int32, error)
	FindByID(params FindOneStoryParams) (StoryResponse, error)
	Like(LikeStoryDTO) (LikeStoryResponse, error)
	Create(storyDTO CreateStoryDTO, scenesDTO []sceneDomain.CreateSceneDTO) (bson.ObjectID, error)
}

type StoryRepository interface {
	Find(filters StoryFilters) ([]StoryResponse, int32, error)
	FindByID(params FindOneStoryParams) (StoryResponse, error)
	StoryExists(id bson.ObjectID) (bool, error)
	Like(dto LikeStoryDTO) (LikeStoryResponse, error)
	Create(dto CreateStoryDTO) (bson.ObjectID, error)
}

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

type FindOneStoryParams struct {
	ID bson.ObjectID `json:"id" bson:"_id"`
	Me bson.ObjectID `json:"me"`
}

// type GetStoryDTO struct {
// 	Data []StoryResponse
// 	Total int
// }

// type StoryResponse struct {
// 	Story  `json:",inline"`
// 	Author struct {
// 		ID    string `bson:"_id" json:"id"`
// 		Login string `bson:"login" json:"login"`
// 	} `bson:"author" json:"author"`
// }
