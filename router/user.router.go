package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, uc controller.UserController) {
	routes := r.Group("/api/user")
	{
		routes.POST("/register", uc.Register)
	}
}
