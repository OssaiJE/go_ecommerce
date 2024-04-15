package services

import (
	"context"
	"go_ecommerce/config"
	"go_ecommerce/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = config.DBCollection(config.ConnectDB(), "users")

func UserExistByEmail(ctx context.Context, email string) (count int64, err error) {
	count, err = UserCollection.CountDocuments(ctx, bson.M{"email": email})

	return count, err
}

func UserExistByPhoneNumber(ctx context.Context, phone_number string) (count int64, err error) {
	count, err = UserCollection.CountDocuments(ctx, bson.M{"phone_number": phone_number})

	return count, err
}

func CreateUser(ctx context.Context, user models.User) (insertedID interface{}, err error) {
	result, err := UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func FindUserByEmail(ctx context.Context, email string) (user *models.User, err error) {
	filter := bson.M{"email": email}
	err = UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(ctx context.Context, user_id primitive.ObjectID, update interface{}) (*models.User, error) {
	filter := bson.M{"_id": user_id}
	_, err := UserCollection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	// Retrieve the updated User
	var updatedUser models.User
	err = UserCollection.FindOne(ctx, filter).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
