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

		userRole, ok := claims["role"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid user Role in token"))
			return
		}

		var deptIdPtr *string
		if deptIdClaim, ok := claims["department_id"].(string); ok && deptIdClaim != "" {
			deptIdPtr = &deptIdClaim
		}

		ctx.Set("user", userId)
		ctx.Set("role", userRole)

		if deptIdPtr != nil {
			ctx.Set("department_id", *deptIdPtr)
		}
		ctx.Next()
	} else {
		tryRefreshToken(ctx)
	}
}

func tryRefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "No access token and refresh token, please login"))
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

	var deptIdPtr *uuid.UUID
	if deptIdClaim, ok := claims["department_id"].(string); ok && deptIdClaim != "" {
		if parsedUUID, err := uuid.Parse(deptIdClaim); err == nil {
			deptIdPtr = &parsedUUID
		}
	}

	newAccessToken, err := utils.GenerateToken(userUUID, claims["role"].(string), deptIdPtr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseFailed(dto.MSG_ACCESS_TOKEN_CREATE_FAILED, "Failed to create access token"))
		return
	}

	ctx.SetCookie("accessToken", newAccessToken, int(time.Minute*120/time.Second), "/", "", false, true)
	ctx.Set("user", userUUID.String())

	if deptIdPtr != nil {
		ctx.Set("department_id", *deptIdPtr)
	}

	ctx.Next()
}

func OnlyAdmin(ctx *gin.Context) {
	userRole, exists := ctx.Get("role")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "No role"))
		return
	}

	if userRole != "admin" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseFailed(dto.MSG_USER_FORBIDDEN, "No permission"))
		return
	}

	ctx.Next()
}
