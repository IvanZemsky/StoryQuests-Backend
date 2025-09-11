package db

import (
	commonErrors "stories-backend/pkg/errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ParseObjectID(hexId string) (bson.ObjectID, error) {
	objID, err := bson.ObjectIDFromHex(hexId)
	if err != nil {
		return bson.ObjectID{}, commonErrors.ErrParsingObjectID
	}
	return objID, nil
}
