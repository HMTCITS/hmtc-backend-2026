package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, uc controller.UserController) {
	routes := r.Group("/api/auth")
	{
		routes.POST("/register", uc.Register)
		routes.GET("/getuser", uc.GetUserByNRP)
		routes.POST("/login", uc.Login)
		routes.POST("/refresh", uc.Refresh)
	}
}
