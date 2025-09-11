package commonErrors

import (
	"errors"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var (
	ErrParsingObjectID = errors.New("object id parsing error")
	ErrNotFound        = mongo.ErrNoDocuments
)