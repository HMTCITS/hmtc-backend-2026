package router

import (
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/HMTCITS/hmtc-backend-2025/middleware"
	"github.com/gin-gonic/gin"
)

func Gallery(r *gin.Engine, gc controller.GalleryController) {
	routes := r.Group("/api/gallery")
	{

		routes.GET("", middleware.RequireAuth, gc.GetAll)
		routes.GET("/:id", middleware.RequireAuth, gc.GetByID)

		routes.POST("", middleware.RequireAuth, middleware.OnlyCMI, gc.Create)
		routes.PUT("/:id", middleware.RequireAuth, middleware.OnlyCMI, gc.Update)
		routes.DELETE("/:id", middleware.RequireAuth, middleware.OnlyCMI, gc.Delete)
	}
}
