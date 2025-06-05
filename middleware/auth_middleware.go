package middleware

import (
	"net/http"
	"strings"

	myjwt "github.com/HMTCITS/hmtc-backend-2025/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func OnlyAllow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRole := ctx.MustGet("role").(string)

		if userRole == "admin" {
			ctx.Next()
			return
		}

		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
	}
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "tidak ada Authorization header"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "format Authorization bearer <token>"})
			return
		}

		tokenString := parts[1]
		payload, err := myjwt.GetPayload(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("id", payload["id"])
		ctx.Set("nrp", payload["nrp"])
		ctx.Set("role", payload["role"])
		ctx.Set("departemen_id", payload["departemen_id"])
		ctx.Set("departemen_name", payload["departemen_name"])

		ctx.Next()
	}
}
