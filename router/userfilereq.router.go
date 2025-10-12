package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func UserFileReq(r *gin.Engine, c controller.UserFileReqController) {
	routes := r.Group("/api/request-file")
	{
		routes.POST("/new-request", c.NewUserFileReq)
		routes.POST("/change-request-status", c.ChangeUserFileReqStatus)
		routes.GET("/:reqid", c.UserFileStatus)
	}
}
