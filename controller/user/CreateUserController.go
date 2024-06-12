package controller

import (
	"context"
	"crudtraining/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	collection *mongo.Collection
}

func NewUserController(db *mongo.Database) *UserController {
	return &UserController{
		collection: db.Collection("users"),
	}
}

func (uc *UserController) CreateUserController(c *gin.Context) {
	var user models.User
	user.ProfilePicture = "profile.jpg"
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" || user.Username == "" || user.Password == "" || user.Nama == "" || user.NoTelp == "" || user.Tanggal_Lahir == "" || user.Jenis_Kelamin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashpassword)
	result, err := uc.collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
