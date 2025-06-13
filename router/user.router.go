package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/HMTCITS/hmtc-backend-2025/middleware"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine, uc controller.UserController) {
	routes := r.Group("/api/auth")
	{
		routes.POST("/register", uc.Register)
		routes.GET("/getuser", uc.GetUserByNRP)
		routes.POST("/login", uc.Login)
		routes.GET("/me", middleware.RequireAuth, uc.Me)
		routes.GET("/admin", middleware.RequireAuth, middleware.OnlyAdmin, uc.MeAdmin)
		routes.POST("/logout", middleware.RequireAuth, uc.Logout)
		routes.POST("/refresh", middleware.RequireAuth, uc.RefreshToken)
	}
}
