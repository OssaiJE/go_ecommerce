package services

import (
	"context"
	"go_ecommerce/config"
	"go_ecommerce/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func FindProductById(ctx context.Context, product_id primitive.ObjectID) (product *models.Product, err error) {
	filter := bson.M{"_id": product_id}
	err = ProductCollection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func FindAllProducts(ctx context.Context) (products []*models.Product, err error) {
	filter := bson.M{}
	cursor, err := ProductCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) 
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}
