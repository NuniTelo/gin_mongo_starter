package models

import (
	"errors"
	"gopkg.in/mgo.v2"
)

/**
This will be used for database connection
 */
var mongoDB *mgo.Database

/**
Here declare all the collections
 */
const(
	post_collection = "posts"
)

/**
SetMongoDB is a function to set a mgo.Database
Set Mongo DatabaseName Here
Also will be used for connection
 */
//
func SetMongoDB(m *mgo.Database) {

	m.Name = "test2" //set the database namemeeeeeee
	mongoDB = m
}

/**
Check if is null
 */
func checkMongoDBNotNull() error {
	if mongoDB == nil {
		return errors.New("mongoDB client is null please set it before use")
	}
	return nil
}
