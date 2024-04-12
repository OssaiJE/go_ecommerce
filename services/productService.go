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

func UpdateOneProduct(ctx context.Context, product_id primitive.ObjectID, user_id primitive.ObjectID, update *models.Product) (*models.Product, error) {
	filter := bson.M{"_id": product_id, "user_id": user_id}
	_, err := ProductCollection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	// Retrieve the updated product
	var updatedProduct models.Product
	err = ProductCollection.FindOne(ctx, filter).Decode(&updatedProduct)
	if err != nil {
		return nil, err
	}

	return &updatedProduct, nil
}

func DeleteOneProduct(ctx context.Context, productID, userID primitive.ObjectID) error {
	filter := bson.M{"_id": productID, "user_id": userID}
	_, err := ProductCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
