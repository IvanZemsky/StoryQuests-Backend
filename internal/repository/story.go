package repository

import (
	"stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type storyRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStoryRepository(db *mongo.Database, collection *mongo.Collection) domain.StoryRepository {
	return &storyRepository{
		db:         db,
		collection: collection,
	}
}

func (repo *storyRepository) Find(filters domain.StoryFilters) ([]domain.Story, error) {
	ctx, cancel := NewCustomRequestTimeoutContext(60)
	defer cancel()

	query := buildFindQuery(&filters)

	cursor, err := repo.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stories []domain.Story
	if err = cursor.All(ctx, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}

func buildFindQuery(filters *domain.StoryFilters) bson.M {
	query := bson.M{}

	if filters.Search != "" {
		query["$or"] = []bson.M{
			{"name": bson.M{"$regex": filters.Search, "$options": "i"}},
			{"description": bson.M{"$regex": filters.Search, "$options": "i"}},
		}
	}

	return query
}

func (repo *storyRepository) FindByID(id bson.ObjectID) (domain.Story, error) {
	ctx, cancel := NewRequestTimeoutContext()
	defer cancel()

	var story domain.Story

	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&story)
	if err != nil {
		return domain.Story{}, err
	}

	return story, nil
}
