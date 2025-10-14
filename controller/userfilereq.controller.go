package controller

import (
	"net/http"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type UserFileReqController interface {
	NewUserFileReq(ctx *gin.Context)
	ChangeUserFileReqStatus(ctx *gin.Context)
	UserFileStatus(ctx *gin.Context)
}

type userFileReqController struct {
	userFileService service.UserFileReqService
}

func NewUserFileReqController(userFileService service.UserFileReqService) UserFileReqController {
	return &userFileReqController{userFileService: userFileService}
}

func (c *userFileReqController) NewUserFileReq(ctx *gin.Context) {

	var req dto.UserFileReqDto
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed("Cannot bind data", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	err := c.userFileService.NewUserFileReq(ctx, req)
	if err != nil {
		res := utils.ResponseFailed("Cannot make request file", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.ResponseSuccess("Request send", nil)
	ctx.JSON(http.StatusAccepted, res)
}

func (c *userFileReqController) ChangeUserFileReqStatus(ctx *gin.Context) {

	var req dto.ChangeUserFileReqDto
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed("Cannot bind data", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := c.userFileService.ChangeUserFileReqStatus(ctx, req)
	if err != nil {
		res := utils.ResponseFailed("Cannot change file request status", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.ResponseSuccess("file status change", nil)
	ctx.JSON(http.StatusAccepted, res)
}

func (c *userFileReqController) UserFileStatus(ctx *gin.Context) {

	reqId := ctx.Param("reqid")

	status, err := c.userFileService.UserFileReqStatus(ctx, reqId)
	if err != nil {
		res := utils.ResponseFailed("Cannot get request file status", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.ResponseSuccess("file status fetched", status)
	ctx.JSON(http.StatusAccepted, res)
}
