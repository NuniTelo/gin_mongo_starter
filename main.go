package main

import (
	"gin-blog/controllers"
	"gin-blog/database"
	"gin-blog/jwt"
	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect() //This will connect to the database

	router := gin.Default() //this will create the gin
	router.Use(database.InjectMongoDB)

	router.Use(jwt.JWT())
	router.GET("/health",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"service": "up",
		})
	})


	api := router.Group("/api/v1")
	{
		posts := api.Group("/posts")
		{
			posts.GET("/all",controllers.SavePost)
		}

	}


	router.Run() // listen and serve on 0.0.0.0:8080


}
