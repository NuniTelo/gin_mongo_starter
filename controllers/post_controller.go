package controllers

import (
	"errors"
	"gin-blog/models"
	"github.com/gin-gonic/gin"
)

func SavePost(c *gin.Context) {

	post_pointer := &models.Post{}

	err:= post_pointer.Save()
	if err != nil {
		c.Error(errors.New("Error on : " + err.Error()))
	}

	c.JSON(200, gin.H{
		"success":true,
	})
}
