package routes

import (
	produkDataController "crudtraining/controller/produk"
	userDataController "crudtraining/controller/user"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()
	db := client.Database("golangpiji_dummyproject")
	userController := userDataController.NewUserController(db)
	produkController := produkDataController.NewProdukController(db)
	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
	})
	//USER SECTION
	users := r.Group("/users")
	{
		users.POST("/usersCreate", userController.CreateUserController)
		users.POST("/usersLogin", userController.UserLoginController)
		users.PUT("/users/update/:id", userController.UpdateUserController)
		users.DELETE("/deleteusers/:id", userController.DeleteUserController)
		users.GET("/users", userController.GetUserController)
		users.GET("/users/:id", userController.GetUserByIdController)
	}
	//PRODUK SECTIOn
	produk := r.Group("/produk")
	{
		produk.POST("/tambahproduk", produkController.CreateProdukController)
	}
	return r
}
