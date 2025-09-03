package handlers

import (
	authDomain "stories-backend/internal/domain/auth"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetUserID(ctx *gin.Context) (bson.ObjectID, error) {
	claims, exists := ctx.Get(authDomain.CTX_AUTH_CLAIMS)
	if !exists {
		return bson.ObjectID{}, nil
	}

	stringUserID := claims.(authDomain.JWTClaims).ID

	userID, err := db.ParseObjectID(stringUserID)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return userID, nil
}