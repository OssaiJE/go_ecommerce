package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Digital bool    `json:"digital" bson:"digital"`
	Cash    bool    `json:"cash" bson:"cash"`
}
