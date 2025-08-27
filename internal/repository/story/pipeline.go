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

var zeroIsLikedPipeline = bson.M{
	"$addFields": bson.M{
		"isLiked": false,
	},
}

func getIsLikedPipeline(me bson.ObjectID) bson.A {
	return bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "stories-likes",
				"localField":   "_id",
				"foreignField": "storyId",
				"as":           "userLikes",
				"pipeline": bson.A{
					bson.M{
						"$match": bson.M{
							"userId": me,
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
	}
}
