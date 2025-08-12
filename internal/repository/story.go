package repository

import (
	"stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

	findOptions := options.Find()

	query := buildFindQuery(&filters)
	sort := buildSortQuery(&filters)

	if len(sort) > 0 {
		findOptions.SetSort(sort)
	}

	cursor, err := repo.collection.Find(ctx, query, findOptions)
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

	if filters.Length != "" {
		if domain.IsValidLength(filters.Length) {
			switch filters.Length {
			case "short":
				query["sceneCount"] = bson.M{
					"$gt": domain.StoryLengthFilterOptions.Short.Gt,
					"$lt": domain.StoryLengthFilterOptions.Short.Lt,
				}
			case "medium":
				query["sceneCount"] = bson.M{
					"$gt": domain.StoryLengthFilterOptions.Medium.Gt,
					"$lt": domain.StoryLengthFilterOptions.Medium.Lt,
				}
			case "long":
				query["sceneCount"] = bson.M{
					"$gt": domain.StoryLengthFilterOptions.Long.Gt,
					"$lt": domain.StoryLengthFilterOptions.Long.Lt,
				}
			}
		}
	}

	return query
}

func buildSortQuery(filters *domain.StoryFilters) bson.D {
	sort := bson.D{}
	if filters.Sort != "" && domain.IsValidSort(filters.Sort) {
		switch filters.Sort {
		case "popular":
			sort = bson.D{{Key: "passes", Value: -1}}
		case "new":
			sort = bson.D{{Key: "date", Value: -1}}
		case "best":
			sort = bson.D{{Key: "likes", Value: -1}}
		}
	}
	return sort
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
