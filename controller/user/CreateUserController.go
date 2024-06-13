package controller

import (
	"context"
	"crudtraining/models"
	"net/http"
	"os"
	"path/filepath"

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
	file, err := c.FormFile("profile_picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar Kosong!"})
		return
	}
	os.MkdirAll("./uploads", os.ModePerm)
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error models": err.Error()})
		return
	}
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashpassword)
	if err != nil {
		c.JSON(400, gin.H{"error password": err.Error()})
		return
	}
	result, err := uc.collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
