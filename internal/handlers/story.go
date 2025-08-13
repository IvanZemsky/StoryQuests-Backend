package handlers

import (
	"errors"
	"net/http"
	"stories-backend/internal/domain/story"
	db "stories-backend/pkg/db/mongo"
	"stories-backend/pkg/errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StoryHandler struct {
	service domain.StoryService
}

func NewStoryHandler(r *gin.Engine, service domain.StoryService) {
	handler := StoryHandler{service: service}

	r.GET("/stories", handler.Find)
	r.GET("/stories/:id", handler.FindByID)
}

func (handler *StoryHandler) Find(ctx *gin.Context) {
	filters, err := handler.parseStoryFilters(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stories, err := handler.service.Find(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, stories)
}

func (handler *StoryHandler) FindByID(ctx *gin.Context) {
	id, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	story, err := handler.service.FindByID(id)

	if err != nil {
		if errors.Is(err, customErrors.ErrParsingObjectID) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, story)
}

func (handler *StoryHandler) parseStoryFilters(ctx *gin.Context) (domain.StoryFilters, error) {
	var filters domain.StoryFilters

	filters.Search = ctx.Query("search")
	filters.Sort = ctx.Query("sort")
	filters.Length = ctx.Query("length")

	limit, err := parseIntQueryParam(ctx, "limit", 0)
	if err != nil {
		return filters, err
	}
	filters.Limit = limit

	page, err := parseIntQueryParam(ctx, "page", 0)
	if err != nil {
		return filters, err
	}
	filters.Page = page

	return filters, nil
}

func parseIntQueryParam(ctx *gin.Context, paramName string, defaultValue int) (int, error) {
	paramStr := ctx.Query(paramName)
	if paramStr == "" {
		return defaultValue, nil
	}

	parsed, err := strconv.Atoi(paramStr)
	if err != nil {
		return 0, errors.New("Invalid " + paramName)
	}
	return parsed, nil
}
