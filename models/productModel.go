package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID           primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	User_ID      primitive.ObjectID `bson:"user_id,omitempty" json:"user_id"`
	Product_Name *string             `json:"product_name,omitempty" bson:"product_name"`
	Category     []*string           `json:"category,omitempty" bson:"category,omitempty"`
	Price        *uint64             `json:"price,omitempty" bson:"price,omitempty"`
	Rating       *uint8              `json:"rating,omitempty" bson:"rating,omitempty"`
	Image        *string             `json:"image,omitempty" bson:"image,omitempty"`
	Created_At   time.Time           `bson:"created_at,omitempty" json:"created_at"`
	Updated_At   time.Time           `bson:"updated_at,omitempty" json:"updated_at"`
}


type ProductUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        *uint64            `json:"price" bson:"price"`
	Rating       *uint8             `json:"rating" bson:"rating"`
	Image        *string            `json:"image" bson:"image"`
}
