package service

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"stories-backend/pkg/errors"
)

func ParseObjectID(hexId string) (bson.ObjectID, error) {
	objID, err := bson.ObjectIDFromHex(hexId)
	if err != nil {
		return bson.ObjectID{}, customErrors.ErrParsingObjectID
	}
	return objID, nil
}
