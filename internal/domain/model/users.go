package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Users struct {
	UUID  bson.ObjectId `json:"uuid" bson:"uuid"`
	Name  string        `json:"name" bson:"name"`
	Email string        `json:"email" bson:"email"`
	Phone int           `json:"phone" bson:"phone"`
	Job   string        `json:"job" bson:"job"`
}
