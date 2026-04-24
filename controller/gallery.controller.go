package controller

import (
	"errors"
	"net/http"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type GalleryController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type galleryController struct {
	galleryService service.GalleryService
	userService    service.UserService
}

func NewGalleryController(gs service.GalleryService, us service.UserService) GalleryController {
	return &galleryController{
		galleryService: gs,
		userService:    us,
	}
}

// Create godoc
// @Summary Create a gallery entry
// @Description Create a new gallery with G-Drive link and metadata (CMI only)
// @Tags gallery
// @Accept json
// @Produce json
// @Param body body dto.GalleryCreateReq true "Gallery data"
// @Success 201 {object} utils.Response{data=dto.GalleryResponse}
// @Failure 400 {object} utils.Response{error=string}
// @Failure 403 {object} utils.Response{error=string}
// @Failure 500 {object} utils.Response{error=string}
// @Router /gallery [post]
func (gc *galleryController) Create(ctx *gin.Context) {
	var req dto.GalleryCreateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseFailed(dto.MSG_GALLERY_CREATE_FAILED, err.Error()))
		return
	}

	result, err := gc.galleryService.CreateGallery(ctx.Request.Context(), req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, dto.ErrGalleryInvalidGDriveLink) ||
			errors.Is(err, dto.ErrGalleryThumbnailRequired) {
			status = http.StatusBadRequest
		}
		ctx.AbortWithStatusJSON(status, utils.ResponseFailed(dto.MSG_GALLERY_CREATE_FAILED, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, utils.ResponseSuccess(dto.MSG_GALLERY_CREATE_SUCCESS, result))
}

// GetAll godoc
// @Summary List galleries
// @Description Get all galleries visible to the authenticated user, with optional date filtering
// @Tags gallery
// @Produce json
// @Param event_date_from query string false "Filter from date (YYYY-MM-DD)"
// @Param event_date_to query string false "Filter to date (YYYY-MM-DD)"
// @Success 200 {object} utils.Response{data=[]dto.GalleryResponse}
// @Failure 500 {object} utils.Response{error=string}
// @Router /gallery [get]
func (gc *galleryController) GetAll(ctx *gin.Context) {
	var filters dto.GalleryFilterParams
	if err := ctx.ShouldBindQuery(&filters); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseFailed(dto.MSG_GALLERY_GET_FAILED, err.Error()))
		return
	}

	result, err := gc.galleryService.GetGalleries(ctx.Request.Context(), filters)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.ResponseFailed(dto.MSG_GALLERY_GET_FAILED, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseSuccess(dto.MSG_GALLERY_GET_SUCCESS, result))
}

// GetByID godoc
// @Summary Get gallery by ID
// @Description Get a single gallery entry by its UUID
// @Tags gallery
// @Produce json
// @Param id path string true "Gallery UUID"
// @Success 200 {object} utils.Response{data=dto.GalleryResponse}
// @Failure 400 {object} utils.Response{error=string}
// @Failure 404 {object} utils.Response{error=string}
// @Router /gallery/{id} [get]
func (gc *galleryController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := gc.galleryService.GetGalleryByID(ctx.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, dto.ErrGalleryInvalidID) {
			status = http.StatusBadRequest
		} else if errors.Is(err, dto.ErrGalleryNotFound) {
			status = http.StatusNotFound
		}
		ctx.AbortWithStatusJSON(status, utils.ResponseFailed(dto.MSG_GALLERY_NOT_FOUND, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseSuccess(dto.MSG_GALLERY_DETAIL_SUCCESS, result))
}

// Update godoc
// @Summary Update a gallery
// @Description Update gallery metadata (CMI only)
// @Tags gallery
// @Accept json
// @Produce json
// @Param id path string true "Gallery UUID"
// @Param body body dto.GalleryUpdateReq true "Updated gallery data"
// @Success 200 {object} utils.Response{data=dto.GalleryResponse}
// @Failure 400 {object} utils.Response{error=string}
// @Failure 403 {object} utils.Response{error=string}
// @Failure 404 {object} utils.Response{error=string}
// @Router /gallery/{id} [put]
func (gc *galleryController) Update(ctx *gin.Context) {
	id := ctx.Param("id")

	var req dto.GalleryUpdateReq
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.ResponseFailed(dto.MSG_GALLERY_UPDATE_FAILED, err.Error()))
		return
	}

	result, err := gc.galleryService.UpdateGallery(ctx.Request.Context(), id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, dto.ErrGalleryInvalidID) ||
			errors.Is(err, dto.ErrGalleryInvalidGDriveLink) ||
			errors.Is(err, dto.ErrGalleryThumbnailRequired) {
			status = http.StatusBadRequest
		} else if errors.Is(err, dto.ErrGalleryNotFound) {
			status = http.StatusNotFound
		}
		ctx.AbortWithStatusJSON(status, utils.ResponseFailed(dto.MSG_GALLERY_UPDATE_FAILED, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseSuccess(dto.MSG_GALLERY_UPDATE_SUCCESS, result))
}

// Delete godoc
// @Summary Delete a gallery
// @Description Delete a gallery entry by ID (CMI only)
// @Tags gallery
// @Produce json
// @Param id path string true "Gallery UUID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response{error=string}
// @Failure 403 {object} utils.Response{error=string}
// @Failure 404 {object} utils.Response{error=string}
// @Router /gallery/{id} [delete]
func (gc *galleryController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := gc.galleryService.DeleteGallery(ctx.Request.Context(), id)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, dto.ErrGalleryInvalidID) {
			status = http.StatusBadRequest
		} else if errors.Is(err, dto.ErrGalleryNotFound) {
			status = http.StatusNotFound
		}
		ctx.AbortWithStatusJSON(status, utils.ResponseFailed(dto.MSG_GALLERY_DELETE_FAILED, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, utils.ResponseSuccess(dto.MSG_GALLERY_DELETE_SUCCESS, nil))
}
