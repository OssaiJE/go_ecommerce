package routes

import (
	"go_ecommerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/auth/signup", controllers.SignUp)
	router.POST("/auth/signin", controllers.SignIn)
	router.GET("/product")
	router.POST("/product/add")
	router.GET("/product/search")
}
