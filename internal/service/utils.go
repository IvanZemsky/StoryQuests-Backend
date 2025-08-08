package service

import "go.mongodb.org/mongo-driver/v2/bson"

func ParseObjectID(hexId string) (bson.ObjectID, error) {
	objID, err := bson.ObjectIDFromHex(hexId)
	if err != nil {
		return bson.ObjectID{}, err
	}
	return objID, nil
}