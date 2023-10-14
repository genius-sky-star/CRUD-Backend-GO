package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	r.POST("/api/login", controllers.Login)
	r.POST("/api/signup", controllers.Signup)
	r.GET("/api/home", controllers.Home)
	r.GET("/api/premium", controllers.Premium)
	r.GET("/api/logout", controllers.Logout)
}
