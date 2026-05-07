package controller

import (
	"net/http"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GalleryController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type galleryController struct {
	galleriesService service.GalleriesService
}

func NewGalleryController(galleriesService service.GalleriesService) GalleryController {
	return &galleryController{
		galleriesService: galleriesService,
	}
}

// Create godoc
// @Summary Buat gallery baru
// @Description Membuat entry gallery baru dengan title, deskripsi, tanggal, link G-Drive, dan thumbnail (CMI Only)
// @Tags gallery
// @Accept json
// @Produce json
// @Param gallery body dto.GalleryCreateReq true "Gallery data"
// @Success 201 {object} utils.Response{data=model.Gallery}
// @Failure 400 {object} utils.Response{error=string} "Bad Request"
// @Failure 500 {object} utils.Response{error=string} "Internal Server Error"
// @Router /gallery [post]
func (c *galleryController) Create(ctx *gin.Context) {
	var req dto.GalleryCreateReq

	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	gallery, err := c.galleriesService.CreateGallery(ctx.Request.Context(), req)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_CREATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_GALLERY_CREATE_SUCCESS, gallery)
	ctx.JSON(http.StatusCreated, res)
}

// GetAll godoc
// @Summary List semua gallery
// @Description Menampilkan list gallery dengan filter berdasarkan hak akses user yang login
// @Tags gallery
// @Produce json
// @Success 200 {object} utils.Response{data=[]model.Gallery}
// @Failure 500 {object} utils.Response{error=string} "Internal Server Error"
// @Router /gallery [get]
func (c *galleryController) GetAll(ctx *gin.Context) {
	userRole, _ := ctx.Get("role")
	userDepartement, _ := ctx.Get("departement")

	role, _ := userRole.(string)
	dept, _ := userDepartement.(string)

	galleries, err := c.galleriesService.GetGalleries(ctx.Request.Context(), role, dept)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_GET_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_GALLERY_GET_SUCCESS, galleries)
	ctx.JSON(http.StatusOK, res)
}

// Update godoc
// @Summary Update gallery
// @Description Update metadata gallery berdasarkan ID (CMI Only)
// @Tags gallery
// @Accept json
// @Produce json
// @Param id path string true "Gallery ID (UUID)"
// @Param gallery body dto.GalleryUpdateReq true "Gallery update data"
// @Success 200 {object} utils.Response{data=model.Gallery}
// @Failure 400 {object} utils.Response{error=string} "Bad Request"
// @Failure 404 {object} utils.Response{error=string} "Not Found"
// @Failure 500 {object} utils.Response{error=string} "Internal Server Error"
// @Router /gallery/{id} [put]
func (c *galleryController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	galleryId, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_INVALID_ID, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var req dto.GalleryUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_UPDATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	gallery, err := c.galleriesService.UpdateGallery(ctx.Request.Context(), galleryId, req)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_UPDATE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_GALLERY_UPDATE_SUCCESS, gallery)
	ctx.JSON(http.StatusOK, res)
}

// Delete godoc
// @Summary Hapus gallery
// @Description Menghapus gallery berdasarkan ID (CMI Only, soft delete)
// @Tags gallery
// @Produce json
// @Param id path string true "Gallery ID (UUID)"
// @Success 200 {object} utils.Response{data=string}
// @Failure 400 {object} utils.Response{error=string} "Bad Request"
// @Failure 404 {object} utils.Response{error=string} "Not Found"
// @Failure 500 {object} utils.Response{error=string} "Internal Server Error"
// @Router /gallery/{id} [delete]
func (c *galleryController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	galleryId, err := uuid.Parse(idParam)
	if err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_INVALID_ID, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := c.galleriesService.DeleteGallery(ctx.Request.Context(), galleryId); err != nil {
		res := utils.ResponseFailed(dto.MSG_GALLERY_DELETE_FAILED, err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess(dto.MSG_GALLERY_DELETE_SUCCESS, nil)
	ctx.JSON(http.StatusOK, res)
}
