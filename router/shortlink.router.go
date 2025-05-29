package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/gin-gonic/gin"
)

func ShortLink(r *gin.Engine, c controller.ShortLinkController) {
	routes := r.Group("/api/shortlink")
	{
		routes.POST("/generate", c.GenerateShortLink)
		routes.GET("/redirect/:shorturl", c.RedirectShortLink)
	}
}
