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
	auth.POST("/signin", middleware.Authenticate, controllers.SignIn)
    
    // Product Routes
	product := router.Group("/product")
	product.GET("/")
	product.POST("/add")
	product.GET("/search")
}
