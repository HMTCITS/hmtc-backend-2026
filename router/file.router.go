package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func FileTA(r *gin.Engine, c controller.FileTAController) {
	routes := r.Group("/file-ta")
	{
		routes.POST("/upload", c.CreateFileTA)
	}
}
