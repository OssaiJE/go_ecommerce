package routes

import (
	"go_ecommerce/controllers"
	"go_ecommerce/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Auth Routes
	auth := router.Group("/auth")
	auth.POST("/signup", controllers.SignUp)
	auth.POST("/signin", controllers.SignIn)

	// User Routes
	user := router.Group("/user")
	user.GET("/", middleware.Authenticate, controllers.GetUser)
	user.PATCH("/update", middleware.Authenticate, controllers.UpdateUser)
	user.PATCH("/photo", middleware.Authenticate, controllers.UpdateProfilePhoto)

	// Product Routes
	product := router.Group("/product")
	product.GET("/", controllers.GetAllProducts)
	product.GET("/:id", controllers.GetOneProduct)
	product.PATCH("/:id", middleware.Authenticate, controllers.UpdateProduct)
	product.DELETE("/:id", middleware.Authenticate, controllers.DeleteProduct)
	product.POST("/create", middleware.Authenticate, controllers.CreateProduct)
	product.GET("/search")
	// product.POST("/uploads", controllers.UploadFiles)
}
