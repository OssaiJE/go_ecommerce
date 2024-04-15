package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_Name   string             `bson:"first_name" json:"first_name" validate:"required,min=2,max=30"`
	Last_Name    string             `bson:"last_name" json:"last_name" validate:"required,min=2,max=30"`
	Email        string             `bson:"email" json:"email" validate:"required,email"`
	Password     string             `bson:"password,omitempty" json:"password" validate:"required,min=6,max=25"`
	Phone_Number *string            `bson:"phone_number" json:"phone_number" validate:"min=9,max=17"`
	// User_Type       string             `bson:"user_type" json:"user_type" validate:"required, eq=admin|eq=user"`
    Image        *string             `json:"image,omitempty" bson:"image,omitempty"`
	User_Cart       []ProductUser `json:"user_cart" bson:"user_cart"`
	Address_Details []Address     `json:"address_details" bson:"address_details"`
	Order_Status    []Order       `json:"order_status" bson:"order_status"`
	Created_At      time.Time     `bson:"created_at,omitempty" json:"created_at"`
	Updated_At      time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}
