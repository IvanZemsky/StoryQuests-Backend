package repository

import (
	"context"
	"errors"
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) Find(filters domain.StoryFilters) ([]domain.StoryResponse, int32, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	stories, totalCount, err := repo.fetchWithCount(ctx, &filters)
	if err != nil {
		return nil, 0, err
	}

	return stories, totalCount, nil
}

func (repo *storyRepository) fetchWithCount(
	ctx context.Context,
	filters *domain.StoryFilters,
) ([]domain.StoryResponse, int32, error) {
	var wg sync.WaitGroup
	var stories []domain.StoryResponse
	var totalCount int32
	var dataErr, countErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		stories, dataErr = repo.fetchStories(ctx, filters)
	}()

	go func() {
		defer wg.Done()
		totalCount, countErr = repo.countStories(ctx, filters)
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

func (repo *storyRepository) fetchStories(
	ctx context.Context,
	filters *domain.StoryFilters,
) ([]domain.StoryResponse, error) {
	pipeline := buildPipeline(filters)

	cursor, err := repo.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stories []domain.StoryResponse
	if err := cursor.All(ctx, &stories); err != nil {
		return nil, err
	}
	return stories, nil
}

func (repo *storyRepository) countStories(ctx context.Context, filters *domain.StoryFilters) (int32, error) {
	pipeline := buildCountPipeline(filters)

	cursor, err := repo.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var countResult []bson.M
	if err := cursor.All(ctx, &countResult); err != nil {
		return 0, err
	}

	if len(countResult) > 0 {
		if total, ok := countResult[0]["total"]; ok {
			if count, ok := total.(int32); ok {
				return count, nil
			}
			return 0, errors.New("field 'total' is not int32")
		}
		return 0, errors.New("field 'total' not found")
	}
	return 0, nil
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
	*pipeline = append(*pipeline, getIsLikedPipeline(filters.Me)...)
}

func addZeroIsLiked(pipeline *bson.A) {
	*pipeline = append(*pipeline, zeroIsLikedPipeline)
}
