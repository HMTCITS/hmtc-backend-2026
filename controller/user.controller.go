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
}

type userController struct {
	userService service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

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
