package database

import (
	connection "gin-blog/models"
	"github.com/gin-gonic/gin"
)

/**
InjectMongoDB is a function that will inject mongo-client to each model
 */
func InjectMongoDB(c *gin.Context) {
	s := Session.Clone()
	c.Set("dbSession", s)
	database := s.DB(Mongo.Database)
	connection.SetMongoDB(database)
	c.Next()
}
