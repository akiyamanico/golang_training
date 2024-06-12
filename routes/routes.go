package routes

import (
	controller "crudtraining/controller/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(client *mongo.Client) *gin.Engine {
	r := gin.Default()
	db := client.Database("golangpiji_dummyproject")
	userController := controller.NewUserController(db)
	r.Use(cors.Default())
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	})
	//USER SECTION
	r.POST("/usersCreate", userController.CreateUserController)
	r.POST("/usersLogin", userController.UserLoginController)
	r.PUT("/users/update/:id", userController.UpdateUserController)
	r.DELETE("/deleteusers/:id", userController.DeleteUserController)
	r.GET("/users", userController.GetUserController)
	r.GET("/users/:id", userController.GetUserByIdController)
	return r
}
