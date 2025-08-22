package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type StoryFilters struct {
	Search string `json:"search"`
	// new | popular | best
	Sort string `json:"sort"`
	// short | medium | long
	Length string `json:"length"`
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	// user setting filters
	Me bson.ObjectID `json:"me"`
}

type GtLtFindOption struct {
	Gt int
	Lt int
}

type StoryLengthFilter struct {
	Short  GtLtFindOption
	Medium GtLtFindOption
	Long   GtLtFindOption
}

var StoryLengthFilterOptions = StoryLengthFilter{
	Short:  GtLtFindOption{Gt: 0, Lt: 10},
	Medium: GtLtFindOption{Gt: 10, Lt: 20},
	Long:   GtLtFindOption{Gt: 20, Lt: 30},
}

func IsValidLength(length string) bool {
	return length == "short" || length == "medium" || length == "long"
}

func IsValidSort(sort string) bool {
	return sort == "popular" || sort == "new" || sort == "best"
}
