package routes

import (
	"github.com/gaisuke/mygram-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoute := r.Group("/users")
	{
		userRoute.POST("/register", func(ctx *gin.Context) {
			controllers.UserRegister(ctx)
		})
		userRoute.POST("/login", func(ctx *gin.Context) {
			controllers.UserLogin(ctx)
		})
	}
}
