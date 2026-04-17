package controller

import (
	"net/http"
	"os"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController interface {
	// Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	// Refresh(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
	GetUserByEmail(ctx *gin.Context)
	Me(ctx *gin.Context)
	MeAdmin(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

// --- COMMENT BENTAR
// // Register godoc
// // @Summary Registrasi user
// // @Description Register user dengan NRP dan asal departemen
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param nrp formData string true "nrp mahasiswa"
// // @Param departement_name formData string true "asal departemen"
// // @Success 201 {object} utils.Response{data=dto.ShortLinkDtoRes}
// // @Failure 400 {object} utils.Response{error=string} "Validation failed or bad request"
// // @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// // @Router /auth/register [post]
// --- COMMENT BENTAR

// func (uc *userController) Register(ctx *gin.Context) {
// 	var userReq dto.UserRegisterReq

// 	if err := ctx.ShouldBind(&userReq); err != nil {
// 		res := utils.ResponseFailed(dto.MSG_USER_REGISTER_FAILED, err.Error())
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
// 		return
// 	}

// 	response, err := uc.userService.Register(userReq)

// 	if err != nil {
// 		res := utils.ResponseFailed(dto.MSG_USER_REGISTER_FAILED, err.Error())
// 		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
// 		return
// 	}

// 	res := utils.ResponseSuccess(dto.MSG_USER_REGISTER_SUCCESS, response)
// 	ctx.JSON(http.StatusCreated, res)
// }

// --- COMMENT BENTAR
// // Get User NRP godoc
// // @Summary Get User NRP
// // @Description Ambil user berdasarkan NRP
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param nrp query string true "Nomor Registrasi Pokok"
// // @Success 200 {object} utils.Response{data=dto.UserGetByNRPRes}
// // @Failure 400 {object} utils.Response{error=string} "Validation failed or NRP kosong"
// // @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// // @Router /auth/getuser [get]
// --- COMMENT BENTAR

func (uc *userController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	if email == "" {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, "Email harus diisi")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userReq := dto.UserGetByEmailReq{Email: email}
	response, err := uc.userService.GetUserByEmail(userReq)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_USER_FOUND, response)
	ctx.JSON(http.StatusOK, res)
}

// --- COMMENT BENTAR
// // Login
// // @Summary Login
// // @Description login dengan nrp
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param nrp formData string true "nrp mahasiswa"
// // @Success 200 {object} utils.Response{data=dto.UserLoginRes}
// // @Failure 400 {object} utils.Response{error=string} "nrp tidak ditemukan"
// // @Failure 400 {object} utils.Response{error=string} "nrp harus diisi"
// // @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// // @Router /auth/login [post]
// --- COMMENT BENTAR

func (uc *userController) Login(ctx *gin.Context) {
	var userReq dto.UserLoginReq

	if err := ctx.ShouldBind(&userReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_LOGIN_FAILED, "Email harus diisi")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := uc.userService.Login(userReq)

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_LOGIN_FAILED, "Email tidak ditemukan")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	ctx.SetCookie("accessToken", response.AccessToken, 3600, "/", "", false, true)
	ctx.SetCookie("refreshToken", response.RefreshToken, 7*24*3600, "/", "", false, true)

	res := utils.ResponseSuccess(dto.MSG_USER_LOGIN_SUCCESS, response)
	ctx.JSON(http.StatusOK, res)
}

// --- COMMENT BENTAR
// // Refresh Token
// // @Summary Refresh token
// // @Description refresh token buat accessToken baru
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param refreshToken cookie string true "refresh token"
// // @Success 200 {object} utils.Response{data=dto.UserRefreshRes}
// // @Failure 400 {object} utils.Response{error=string} "Refresh token not found"
// // @Failure 400 {object} utils.Response{error=string} "Invalid refresh token"
// // @Failure 400 {object} utils.Response{error=string} "unauthorized"
// // @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// // @Router /auth/refresh [post]
// --- COMMENT BENTAR

func (uc *userController) RefreshToken(ctx *gin.Context) {
	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
		return
	}

	claims, err := utils.VerifyToken(refreshToken, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	userIDStr, ok := claims["sub"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userUUID, _ := uuid.Parse(userIDStr)

	userDepartement, ok := claims["departement"].(string)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// generate token baru
	newAccessToken, _ := utils.GenerateToken(userUUID, claims["role"].(string), userDepartement)

	ctx.SetCookie("accessToken", newAccessToken, 3600, "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Access token refreshed"})
}

// --- COMMENT BENTAR
// // console me
// // @Summary show user info
// // @Description show user info nrp, departemen
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param user cookie string true "user"
// // @Success 200 {object} utils.Response{data=dto.UserMeRes}
// // @Failure 400 {object} utils.Response{error=string} "User not found"
// // @Failure 400 {object} utils.Response{error=string} "unauthorized"
// // @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// // @Router /auth/me [get]
// --- COMMENT BENTAR

func (uc *userController) Me(ctx *gin.Context) {
	userID, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	response, err := uc.userService.Me(userID.(string))
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_USER_FOUND, response)
	ctx.JSON(http.StatusOK, res)
}

// --- COMMENT BENTAR
// // console admin
// // @Summary show admin info (only for admin)
// // @Description show admin info nrp, departemen
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param user cookie string true "user"
// // @Success 200 {object} utils.Response{data=dto.UserMeRes}
// // @Failure 400 {object} utils.Response{error=string} "User not found"
// // @Failure 400 {object} utils.Response{error=string} "unauthorized"
// // @Failure 400 {object} utils.Response{error=string} "forbidden"
// // @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// // @Router /auth/admin [get]
// --- COMMENT BENTAR

func (uc *userController) MeAdmin(ctx *gin.Context) {
	userID, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	response, err := uc.userService.Me(userID.(string))
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_USER_FOUND, response)
	ctx.JSON(http.StatusOK, res)
}

// --- COMMENT BENTAR
// // Logout
// // @Summary logout
// // @Description logout
// // @Tags user
// // @Accept x-www-form-urlencoded
// // @Produce json
// // @Param accessToken cookie string true "access token"
// // @Param refreshToken cookie string true "refresh token"
// // @Success 200 {object} utils.Response
// // @Router /auth/logout [post]
// --- COMMENT BENTAR

func (uc *userController) Logout(ctx *gin.Context) {
	ctx.SetCookie("accessToken", "", -1, "/", "", false, true)
	ctx.SetCookie("refreshToken", "", -1, "/", "", false, true)

	res := utils.ResponseSuccess("Logout berhasil", nil)
	ctx.JSON(http.StatusOK, res)
}
