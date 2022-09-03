package db

import (
	"context"
	"time"

	"github.com/jjrb3/go-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter_example")
	coll := db.Collection("users")

	u.Password, _ = encrypt(u.Password)

	res, err := coll.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := res.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

func FindByEmail(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter_example")
	coll := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := coll.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
