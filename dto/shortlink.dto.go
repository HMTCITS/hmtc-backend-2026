package dto

type ShortLinkDto struct {
	Link string `json:"link" form:"link" binding:"required"`
}
