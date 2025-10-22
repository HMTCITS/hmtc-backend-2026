package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/HMTCITS/hmtc-backend-2025/middleware"
	"github.com/gin-gonic/gin"
)

func Magang(r *gin.Engine, mc controller.MagangController) {
	routes := r.Group("/api/magang")
	{
		routes.GET("/get-token", mc.GetToken)
		routes.GET("/oauth2callback", mc.Callback)
		// protect upload with schedule middleware (explicit path provided)
		routes.POST("/upload", middleware.RequireScheduleUpload("/api/magang/upload"), mc.Upload)
	}
}
