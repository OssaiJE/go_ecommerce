package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	House       *string            `json:"house" bson:"house"`
	Street      *string            `json:"street" bson:"street"`
	City        *string            `json:"city" bson:"city"`
	Postal_Code *string            `json:"postal_code" bson:"postal_code"`
}
