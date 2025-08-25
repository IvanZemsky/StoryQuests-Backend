package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) Find(filters domain.StoryFilters) ([]domain.StoryResponse, int32, error) {
	ctx, cancel := repository.NewCustomRequestTimeoutContext(60)
	defer cancel()

	// Создаем pipeline для данных
	dataPipeline := buildPipeline(&filters)

	// Создаем pipeline для подсчета (без пагинации и сортировки)
	countPipeline := buildCountPipeline(&filters)

	// Выполняем оба запроса параллельно
	var wg sync.WaitGroup
	var stories []domain.StoryResponse
	var totalCount int32
	var dataErr, countErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		cursor, err := repo.collection.Aggregate(ctx, dataPipeline)
		if err != nil {
			dataErr = err
			return
		}
		defer cursor.Close(ctx)
		dataErr = cursor.All(ctx, &stories)
	}()

	go func() {
		defer wg.Done()
		cursor, err := repo.collection.Aggregate(ctx, countPipeline)
		if err != nil {
			countErr = err
			return
		}
		defer cursor.Close(ctx)

		var countResult []bson.M
		if err := cursor.All(ctx, &countResult); err != nil {
			countErr = err
			return
		}

		if len(countResult) > 0 {
			totalCount = countResult[0]["total"].(int32)
		}
	}()

	wg.Wait()

	if dataErr != nil {
		return nil, 0, dataErr
	}
	if countErr != nil {
		return nil, 0, countErr
	}

	return stories, totalCount, nil
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

func buildCountPipeline(filters *domain.StoryFilters) bson.A {
	query := buildFindQuery(filters)

	pipeline := bson.A{
		bson.M{"$match": query},
	}

	pipeline = append(pipeline,
		bson.M{"$count": "total"},
	)

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
