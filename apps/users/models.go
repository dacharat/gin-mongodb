package users

import "go.mongodb.org/mongo-driver/bson"

type User struct {
	Id     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
}
