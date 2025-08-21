package repository

import "go.mongodb.org/mongo-driver/v2/bson"

var authorPipelineWithoutMatch = bson.A{
	bson.M{
		"$lookup": bson.M{
			"from":         "users",
			"localField":   "author",
			"foreignField": "_id",
			"as":           "author",
		},
	},
	bson.M{"$unwind": "$author"},
	bson.M{
		"$addFields": bson.M{
			"author": bson.M{
				"_id":   "$author._id",
				"login": "$author.login",
			},
		},
	},
}
