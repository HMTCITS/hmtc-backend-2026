package controller

import (
	"net/http"
	"os"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type FileTAController interface {
	CreateFileTA(ctx *gin.Context)
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

	if err := c.fileService.CreateFileTA(ctx, req, file.Filename); err != nil {
		res := utils.ResponseFailed(dto.FileNotCreate, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	os.MkdirAll("./file-ta", os.ModePerm)

	ctx.SaveUploadedFile(file, "./file-ta/"+file.Filename)

	ctx.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
