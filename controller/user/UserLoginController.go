package controller

import (
	"crudtraining/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

var jwtKey = []byte("eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW4iLCJJc3N1ZXIiOiJJc3N1ZXIiLCJVc2VybmFtZSI6IkphdmFJblVzZSIsImV4cCI6MTcxODE4NDIxOSwiaWF0IjoxNzE4MTg0MjE5fQ.-I4tTC7ZbMUVHE1RZnn3Rboi5G0BZupEH70xVj1DwZo")

type Claims struct {
	Username string `json:"Username"`
	jwt.StandardClaims
}

func (uc *UserController) UserLoginController(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := uc.collection.FindOne(context.Background(), bson.M{"username": loginRequest.Username}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau Password Yang Anda Masukan Salah!"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email atau Password Yang Anda Masukan Salah!"})
		return
	}

	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token not signed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Welcome!", "tokenString": tokenString})

}
