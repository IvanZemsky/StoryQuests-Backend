package handlers

import (
	"log"
	"net/http"
	domain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/common"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// @Summary Find stories
// @Description Retrieves stories with optional filtering and pagination (default length is 10)
// @Tags Story
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param sort query string false "Sort field" default(created_at)
// @Param order query string false "Sort order" Enums(asc, desc) default(desc)
// @Param search query string false "Search query"
// @Param category query string false "Category filter"
// @Param status query string false "Status filter" Enums(active, inactive, draft)
// @Success 200 {array} domain.Story "Stories retrieved successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid filter parameters"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Router /stories [get]
func (handler *StoryHandler) Find(ctx *gin.Context) {
	filters, err := handler.parseStoryFilters(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stories, count, err := handler.service.Find(filters)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.Header("X-Total-Count", handlers.Int32ToString(count))
	ctx.JSON(http.StatusOK, stories)
}

func (handler *StoryHandler) parseStoryFilters(ctx *gin.Context) (domain.StoryFilters, error) {
	var filters domain.StoryFilters

	filters.Search = ctx.Query("search")
	filters.Sort = ctx.Query("sort")
	filters.Length = ctx.Query("length")

	authClaims, err := handlers.GetAuthClaims(ctx)
	if err != nil {
		filters.Me = bson.NilObjectID
	} else {
		meID, err := db.ParseObjectID(authClaims.ID)
		if err != nil {
			filters.Me = bson.NilObjectID
		} else {
			filters.Me = meID
		}
	}

	limit, err := handlers.ParseIntQueryParam(ctx.Query("limit"), "limit", 10)
	if err != nil {
		return filters, err
	}
	filters.Limit = limit

	page, err := handlers.ParseIntQueryParam(ctx.Query("page"), "page", 1)
	if err != nil {
		return filters, err
	}
	filters.Page = page

	byUserID, err := db.ParseObjectID(ctx.Query("byUserId"))
	if err != nil {
		filters.ByUserID = bson.NilObjectID
	} else {
		filters.ByUserID = byUserID
	}

	return filters, nil
}
