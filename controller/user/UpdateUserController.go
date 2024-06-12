package controller

import (
	"context"
	"crudtraining/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *UserController) UpdateUserController(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error models": err.Error()})
		return
	}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "username", Value: user.Username},
			{Key: "email", Value: user.Email},
			{Key: "password", Value: user.Password},
		}},
	}
	result, err := uc.collection.UpdateByID(context.TODO(), id, update)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
