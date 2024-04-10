package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Order_Cart     []ProductUser      `json:"order_cart" bson:"order_cart"`
	Ordered_At     time.Time          `bson:"created_at,omitempty" json:"created_at"`
	Price          *uint64            `json:"price" bson:"price"`
	Discount       *int               `json:"dicount" bson:"dicount"`
	Payment_Method Payment            `json:"payment_method" bson:"payment_method"`
}
