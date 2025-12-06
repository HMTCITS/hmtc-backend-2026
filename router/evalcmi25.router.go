package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func EvalCmi25(r *gin.Engine, ecc controller.EvalCmi25Controller) {
	routes := r.Group("/api/evaluasi-cmi")
	{
		routes.POST("/submit", ecc.Upload)
	}
}
