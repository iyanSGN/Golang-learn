package service

import (
	"rearrange/app/kabupaten"
	"rearrange/app/kabupaten/repository"
	"rearrange/package/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB *gorm.DB
	Repository repository.Repository
}

func NewService (DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl)GetAll(c echo.Context) ([]kabupaten.KabKotaResponseDTO, error) {
	var kabupatenRes []kabupaten.KabKotaResponseDTO
	
	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return kabupatenRes, response.BuildError(response.ErrServerError, err)
	}

	for _, kabupaten := range result {
		kabupatenRes = append(kabupatenRes, kabupaten.ToResponse())
	}

	return kabupatenRes, nil

}

func (s *serviceImpl)GetByID(c echo.Context, ID uint) (kabupaten.KabKotaResponseDTO, error) {
	var kabupaten kabupaten.KabKotaResponseDTO

	result, err := s.Repository.GetByID(c, s.DB, ID) 
	if err != nil {
		return kabupaten, response.BuildError(response.ErrNotFound, err)
	}

	kabupaten = result.ToResponse()

	return kabupaten, nil
}