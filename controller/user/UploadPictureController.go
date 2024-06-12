package controller

import (
	"crudtraining/models"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func (uc *UserController) UploadPictureController(c *gin.Context) {
	var user models.User
	file, err := c.FormFile("picture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar Kosong!"})
		return
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ini Bukan File Gambar!"})
		return
	}
	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	filePath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	filter := bson.M{"username": user.Username}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "picture", Value: file.Filename},
		}},
	}
	_, err = uc.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"picture": file.Filename})

}
