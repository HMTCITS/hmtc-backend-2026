package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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

		userDepartement, ok := claims["department"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "Invalid user Departement in token"))
			return
		}

		ctx.Set("user", userId)
		ctx.Set("role", userRole)
		ctx.Set("departement", userDepartement)
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

	newAccessToken, err := utils.GenerateToken(userUUID, claims["role"].(string), claims["department"].(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseFailed(dto.MSG_ACCESS_TOKEN_CREATE_FAILED, "Failed to create access token"))
		return
	}

	ctx.SetCookie("accessToken", newAccessToken, int(24*time.Hour/time.Second), "/", "", false, true)
	ctx.Set("user", userUUID.String())

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

func OnlyCMI(ctx *gin.Context) {
	userDepartment, exists := ctx.Get("departement")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusForbidden, utils.ResponseFailed(dto.MSG_AUTH_FAILED, "No departement"))
		return
	}
	deprt := strings.ToLower(userDepartment.(string))
	if deprt != "cmi" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ResponseFailed(dto.MSG_USER_FORBIDDEN, "Departement has no access"))
		return
	}
	ctx.Next()
}
