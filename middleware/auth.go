package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("accessToken")
	if err != nil {
		tryRefreshToken(ctx)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		tryRefreshToken(ctx)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, ok := claims["sub"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid user ID in token"))
			return
		}
		ctx.Set("user", userId)
		ctx.Next()
	} else {
		tryRefreshToken(ctx)
	}
}

func tryRefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "No refresh token"))
		return
	}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
	})
	if err != nil || !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_INVALID_TOKEN_FAILED, "Invalid refresh token"))
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid refresh token claims"))
		return
	}

	userIdStr, ok := claims["sub"].(string)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid user ID in refresh token"))
		return
	}

	userUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid UUID"))
		return
	}

	// Generate new access token
	newAccessToken, err := utils.GenerateToken(userUUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseFailed(dto.MSG_ACCESS_TOKEN_CREATE_FAILED, "Failed to create access token"))
		return
	}

	// Set new access token
	ctx.SetCookie("accessToken", newAccessToken, int(time.Minute*120/time.Second), "/", "", false, true)
	ctx.Set("user", userUUID.String())
	ctx.Next()
}
