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

func (c *shortLinkController) GenerateShortLink(ctx *gin.Context) {
	var link dto.ShortLinkDto

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
