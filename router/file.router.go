package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/HMTCITS/hmtc-backend-2025/middleware"
	"github.com/gin-gonic/gin"
)

func FileTA(r *gin.Engine, c controller.FileTAController) {
	routes := r.Group("/api/file-ta")
	{
		routes.POST("/upload", c.CreateFileTA)
		routes.POST("/change-status", c.ChangeFileStatus)
		routes.GET("/", c.GetAllFiles)
		routes.GET("/:fileid", c.GetFileStatus)
		routes.GET("/download/:reqid/:fileid", middleware.RequireAuth, c.DownloadFile)
	}
}
