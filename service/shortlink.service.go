package service

import (
	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
)

type ShortLinkService interface {
	GenerateShortLink(link dto.ShortLinkDtoReq) (dto.ShortLinkDtoRes, error)
	FindByShortUrl(link string) (model.LinkShortener, error)
}

type shortLinkService struct {
	repo repository.ShortLinkRepository
}

func NewShortLinkService(repo repository.ShortLinkRepository) ShortLinkService {
	return &shortLinkService{repo: repo}
}

func (s *shortLinkService) GenerateShortLink(request dto.ShortLinkDtoReq) (dto.ShortLinkDtoRes, error) {
	newLink := model.LinkShortener{
		Fullurl:  request.Link,
		Shorturl: request.ShortLink,
	}

	createdLink, err := s.repo.GenerateShortLink(newLink)
	if err != nil {
		return dto.ShortLinkDtoRes{ShortLink: ""}, err
	}

	response := dto.ShortLinkDtoRes{
		ShortLink: "localhost:3000/api/shortlink/redirect/" + createdLink.Shorturl,
	}

	return response, nil
}

func (s *shortLinkService) FindByShortUrl(shortUrl string) (model.LinkShortener, error) {
	return s.repo.FindByShortUrl(shortUrl)
}
