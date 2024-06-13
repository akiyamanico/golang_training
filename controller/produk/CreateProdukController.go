package controller

import (
	"crudtraining/models"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProdukController struct {
	collection *mongo.Collection
}

func NewProdukController(db *mongo.Database) *ProdukController {
	return &ProdukController{
		collection: db.Collection("produk"),
	}
}

func (pc *ProdukController) CreateProdukController(c *gin.Context) {
	var produk models.Produk
	file, err := c.FormFile("gambar_produk")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar Kosong!"})
		return
	}
	os.MkdirAll("./uploads/produk", os.ModePerm)
	filepath.Join("uploads/produk", file.Filename)
	if err := c.ShouldBind(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := pc.collection.InsertOne(c, produk)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}
