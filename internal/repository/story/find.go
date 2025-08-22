package repository

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
)

func (repo *storyRepository) Find(filters domain.StoryFilters) ([]domain.StoryResponse, error) {
	ctx, cancel := repository.NewCustomRequestTimeoutContext(60)
	defer cancel()

	pipeline := buildPipeline(&filters)

	cursor, err := repo.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stories []domain.StoryResponse
	if err = cursor.All(ctx, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}

func buildPipeline(filters *domain.StoryFilters) bson.A {
	query := buildFindQuery(filters)

	pipeline := bson.A{
		bson.M{"$match": query},
	}

	if !filters.Me.IsZero() {
		addIsLiked(&pipeline, filters)
	} else {
		addZeroIsLiked(&pipeline)
	}

	pipeline = append(pipeline, authorPipelineWithoutMatch...)

	sort := buildSortQuery(filters)
	if len(sort) > 0 {
		pipeline = append(pipeline, bson.M{"$sort": sort})
	}

	addPagination(&pipeline, filters.Page, filters.Limit)

	return pipeline
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

func addPagination(pipeline *bson.A, page int, limit int) {
	if page > 0 && limit > 0 {
		*pipeline = append(*pipeline,
			bson.M{"$skip": (page - 1) * limit},
			bson.M{"$limit": limit},
		)
	}
}

func addIsLiked(pipeline *bson.A, filters *domain.StoryFilters) {
	*pipeline = append(*pipeline,
		bson.M{
			"$lookup": bson.M{
				"from":         "stories-likes",
				"localField":   "_id",
				"foreignField": "storyId",
				"as":           "userLikes",
				"pipeline": bson.A{
					bson.M{
						"$match": bson.M{
							"userId": &filters.Me,
						},
					},
				},
			},
		},
		bson.M{
			"$addFields": bson.M{
				"isLiked": bson.M{
					"$gt": bson.A{"$userLikes", []any{}},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"userLikes": 0,
			},
		},
	)
}

func addZeroIsLiked(pipeline *bson.A) {
	*pipeline = append(*pipeline,
		bson.M{
			"$addFields": bson.M{
				"isLiked": false,
			},
		},
	)
}
