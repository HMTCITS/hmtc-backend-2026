package controller

import (
	"net/http"

	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/HMTCITS/hmtc-backend-2025/utils"
	"github.com/gin-gonic/gin"
)

type ShortLinkController interface {
	GenerateShortLink(ctx *gin.Context)
	RedirectShortLink(ctx *gin.Context)
}

type shortLinkController struct {
	shortLinkService service.ShortLinkService
}

func NewShortLinkController(service service.ShortLinkService) ShortLinkController {
	return &shortLinkController{
		shortLinkService: service,
	}
}

// Generate short link godoc
// @Summary buat short link
// @Description buat short link
// @Tags link shortener
// @Accept x-www-form-urlencoded
// @Produce json
// @Param link formData string true "URL asli yang ingin diperpendek"
// @Param shortlink formData string true "nama url (contoh: link:myanimelist.net shortlink: anime)"
// @Success 200 {object} utils.Response{data=dto.ShortLinkDtoRes}
// @Failure 400 {object} utils.Response{error=string} "Bad Request"
// @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// @Router /shortlink/generate [post]
func (c *shortLinkController) GenerateShortLink(ctx *gin.Context) {
	var link dto.ShortLinkDtoReq

	if err := ctx.ShouldBind(&link); err != nil {
		res := utils.ResponseFailed("Invalid link", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	createdLink, err := c.shortLinkService.GenerateShortLink(link)
	if err != nil {
		res := utils.ResponseFailed("Failed to generate short link", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
		return
	}

	res := utils.ResponseSuccess("Short link generated successfully", createdLink)
	ctx.JSON(http.StatusOK, res)
}

// Redirect godoc
// @Summary Redirect ke url asli
// @Description Redirect ke url asli berdasarkan shortlink
// @Tags link shortener
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param shorturl path string true "nama link yang diinginkan"
// @Success 301 {string} string "Redirected"
// @Failure 400 {object} utils.Response{error=string} "Bad Request"
// @Failure 500 {object} utils.Response{error=string} "Internal server or database error"
// @Router /shortlink/redirect/{shorturl} [get]
func (c *shortLinkController) RedirectShortLink(ctx *gin.Context) {
	shortUrl := ctx.Param("shorturl")

	link, err := c.shortLinkService.FindByShortUrl(shortUrl)
	if err != nil {
		res := utils.ResponseFailed("shortlink not found", err.Error())
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, link.Fullurl)
}
