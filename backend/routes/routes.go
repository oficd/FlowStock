package routes

import (
	"backend/controllers"
	"backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	//initialize gin
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// route register
	router.POST("/api/register", controllers.Register)

	// route login
	router.POST("/api/login", controllers.Login)



	// route users
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)

	// route user create
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)

	// route user by id
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUserById)

	// route user update
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)

	// route user delete
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)



	// route products
	router.GET("/api/categories", middlewares.AuthMiddleware(), controllers.FindCategory)

	// route product by id
	router.GET("/api/categories/:id", middlewares.AuthMiddleware(), controllers.FindCategoryById)

	// route product create
	router.POST("/api/categories", middlewares.AuthMiddleware(), controllers.CreateCategory)

	// route product update
	router.PUT("/api/categories/:id", middlewares.AuthMiddleware(), controllers.UpdateCategory)

	// route product delete
	router.DELETE("/api/categories/:id", middlewares.AuthMiddleware(), controllers.DeleteCategory)


	
	// route products
	router.GET("/api/products", middlewares.AuthMiddleware(), controllers.FindProduct)

	// route product by id
	router.GET("/api/products/:id", middlewares.AuthMiddleware(), controllers.FindProductById)

	// route product create
	router.POST("/api/products", middlewares.AuthMiddleware(), controllers.CreateProduct)

	// route product update
	router.PUT("/api/products/:id", middlewares.AuthMiddleware(), controllers.UpdateProduct)

	// route product delete
	router.DELETE("/api/products/:id", middlewares.AuthMiddleware(), controllers.DeleteProduct)



	// route products
	router.GET("/api/suppliers", middlewares.AuthMiddleware(), controllers.FindSupplier)

	// route product by id
	router.GET("/api/suppliers/:id", middlewares.AuthMiddleware(), controllers.FindSupplierById)

	// route product create
	router.POST("/api/suppliers", middlewares.AuthMiddleware(), controllers.CreateSupplier)

	// route product update
	router.PUT("/api/suppliers/:id", middlewares.AuthMiddleware(), controllers.UpdateSupplier)

	// route product delete
	router.DELETE("/api/suppliers/:id", middlewares.AuthMiddleware(), controllers.DeleteSupplier)



	return router
}
