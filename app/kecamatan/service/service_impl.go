package service

import (
	"rearrange/app/kecamatan"
	"rearrange/app/kecamatan/repository"
	"rearrange/package/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB 				*gorm.DB
	Repository		repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB :		 DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) GetAll(c echo.Context) ([]kecamatan.KecamatanResponseDTO, error) {
	var kecamatanRes []kecamatan.KecamatanResponseDTO

	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return kecamatanRes, response.BuildError(response.ErrServerError, err)
	}

	for _, kecamatan := range result {
		kecamatanRes = append(kecamatanRes, kecamatan.ToResponse())
	}

	return kecamatanRes, nil
}

func (s *serviceImpl) GetByID(c echo.Context, ID uint) (kecamatan.KecamatanResponseDTO, error) {
	var kecamatan kecamatan.KecamatanResponseDTO

	result, err := s.Repository.GetByID(c, s.DB, ID)
	if err != nil {
		return kecamatan, response.BuildError(response.ErrNotFound, err)
	}

	kecamatan = result.ToResponse()

	return kecamatan,nil
}