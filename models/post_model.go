package models

import (
	_ "errors"
	_ "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `bson:"name" json:"name"`
}


type Posts[] Post

/**
Save function for topic
 */
func (t *Post) Save() error {
	err := checkMongoDBNotNull()
	err = mongoDB.C(post_collection).Insert(t)
	return err
}

