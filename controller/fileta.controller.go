package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type FileTAController interface {
	CreateFileTA(ctx *gin.Context)
	DownloadFile(ctx *gin.Context)
	ChangeFileStatus(ctx *gin.Context)
	GetAllFiles(ctx *gin.Context)
	GetFileStatus(ctx *gin.Context)
}

type fileTAController struct {
	fileService service.FileTAService
}

func NewFileTAController(fileTAservice service.FileTAService) FileTAController {
	return &fileTAController{fileService: fileTAservice}
}

func (c *fileTAController) CreateFileTA(ctx *gin.Context) {

	var req dto.CreateFileTA
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed("cannot bind data", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		res := utils.ResponseFailed("file is required", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	f, err := file.Open()
	if err != nil {
		res := utils.ResponseFailed("failed to open file", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	defer f.Close()

	buf := make([]byte, 4)
	if _, err := f.Read(buf); err != nil {
		res := utils.ResponseFailed("failed to read file", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if string(buf) != "%PDF" {
		res := utils.ResponseFailed("invalid file type", "file must be a PDF")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if file.Size > 10*1024*1024 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "file too large"})
		return
	}

	if err := c.fileService.CreateFileTA(ctx, req, file.Filename, f); err != nil {
		res := utils.ResponseFailed(dto.FileNotCreate, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func (c *fileTAController) GetFileStatus(ctx *gin.Context) {
	fileId := ctx.Param("fileid")

	status, err := c.fileService.GetFileStatus(ctx, fileId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get file status"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"status": status})
}

func (c *fileTAController) ChangeFileStatus(ctx *gin.Context) {

	var req dto.ChangeFileStatus
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed("cannot bind data", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	err := c.fileService.ChangeFileStatus(ctx, req)
	if err != nil {
		res := utils.ResponseFailed("cannot change status", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess("file status changed", nil)
	ctx.JSON(http.StatusOK, res)
}

func (c *fileTAController) GetAllFiles(ctx *gin.Context) {

	fileTa, err := c.fileService.GetAllFiles(ctx)
	if err != nil {
		res := utils.ResponseFailed(dto.FileNotCreate, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	res := utils.ResponseSuccess("all file fetched", fileTa)
	ctx.JSON(http.StatusAccepted, res)
}

func (c *fileTAController) DownloadFile(ctx *gin.Context) {
	reqId := ctx.Param("reqid")
	fileId := ctx.Param("fileid")
	userNrp, exists := ctx.Get("nrp")
	if !exists {
		res := utils.ResponseFailed("NRP not provided, please login", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	nrpStr := fmt.Sprintf("%v", userNrp)

	filename, err := c.fileService.GetFileName(ctx, reqId, fileId, nrpStr)
	if err != nil {
		res := utils.ResponseFailed("Cannot download file", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}
	path := "./file-ta/" + filename

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.FileAttachment(path, filename)

}
