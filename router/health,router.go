package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.Engine, hc controller.HealthController) {
	r.GET("/health", hc.CheckHealth)
}
