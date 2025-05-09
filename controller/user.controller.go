package controller

import (
	"net/http"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(ctx *gin.Context)
	GetUserByNRP(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

// Register godoc
// @Summary Registrasi user
// @Description Register user dengan NRP dan asal departemen
// @Tags user
// @Accept x-www-form-urlencoded
// @Produce json
// @Param nrp formData string true "nrp mahasiswa"
// @Param departement_name formData string true "asal departemen"
// @Success 201 {object} utils.Response{data=dto.ShortLinkDtoRes}
// @Failure 400 {object} utils.Response{error=string} "Validation failed or bad request"
// @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// @Router /user/register [post]
func (uc *userController) Register(ctx *gin.Context) {
	var userReq dto.UserRegisterReq

	if err := ctx.ShouldBind(&userReq); err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_REGISTER_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	response, err := uc.userService.Register(userReq)

	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_REGISTER_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_USER_REGISTER_SUCCESS, response)
	ctx.JSON(http.StatusCreated, res)
}

// Get User NRP godoc
// @Summary Get User NRP
// @Description Ambil user berdasarkan NRP
// @Tags user
// @Accept x-www-form-urlencoded
// @Produce json
// @Param nrp query string true "Nomor Registrasi Pokok"
// @Success 200 {object} utils.Response{data=dto.UserGetByNRPRes}
// @Failure 400 {object} utils.Response{error=string} "Validation failed or NRP kosong"
// @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// @Router /user/getuser [get]
func (uc *userController) GetUserByNRP(ctx *gin.Context) {
	nrp := ctx.Query("nrp")

	if nrp == "" {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, "NRP harus diisi")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	userReq := dto.UserGetByNRPReq{NRP: nrp}
	response, err := uc.userService.GetUserByNRP(userReq)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_USER_NOT_FOUND, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_USER_FOUND, response)
	ctx.JSON(http.StatusOK, res)
}
