package services

import (
	"context"
	"go_ecommerce/config"
	"go_ecommerce/models"

	"go.mongodb.org/mongo-driver/mongo"
)

var ProductCollection *mongo.Collection = config.DBCollection(config.ConnectDB(), "products")

func CreateProduct(ctx context.Context, product models.Product) (insertedID interface{}, err error) {
	result, err := ProductCollection.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}
