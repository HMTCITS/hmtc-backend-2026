package dto

type ShortLinkDtoReq struct {
	Link      string `json:"link" form:"link" binding:"required"`
	ShortLink string `json:"shortlink" form:"shortlink" binding:"required"`
}

type ShortLinkDtoRes struct {
	ShortLink string `json:"shortlink"`
}
