package controller

import (
	"net/http"

	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type HealthController interface {
	CheckHealth(ctx *gin.Context)
}

type healthController struct{}

func NewHealthController() HealthController {
	return &healthController{}
}

func (hc *healthController) CheckHealth(ctx *gin.Context) {
	healthStatus := map[string]string{
		"status": "OK",
		"env":    "production",
	}

	res := utils.ResponseSuccess("Health check successful", healthStatus)
	ctx.JSON(http.StatusOK, res)
}
