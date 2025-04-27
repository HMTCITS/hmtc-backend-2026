package service

import (
	"github.com/HMTCITS/hmtc-backend-2025/dto"
	"github.com/HMTCITS/hmtc-backend-2025/model"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
)

type ShortLinkService interface {
	GenerateShortLink(link dto.ShortLinkDto) (dto.ShortLinkDto, error)
	FindByShortUrl(link string) (model.LinkShortener, error)
}

type shortLinkService struct {
	repo repository.ShortLinkRepository
}

func NewShortLinkService(repo repository.ShortLinkRepository) ShortLinkService {
	return &shortLinkService{repo: repo}
}

func (s *shortLinkService) GenerateShortLink(request dto.ShortLinkDto) (dto.ShortLinkDto, error) {
	newLink := model.LinkShortener{
		Fullurl: request.Link,
	}

	createdLink, err := s.repo.GenerateShortLink(newLink)
	if err != nil {
		return dto.ShortLinkDto{}, err
	}

	response := dto.ShortLinkDto{
		Link: createdLink.Shorturl, // balikin shortlinknya tapi masih uuidnya aja
	}

	return response, nil
}

func (s *shortLinkService) FindByShortUrl(shortUrl string) (model.LinkShortener, error) {
	return s.repo.FindByShortUrl(shortUrl)
}
