package service

import (
	"rearrange/app/biostar"
	"rearrange/app/biostar/repository"


	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceBio struct {
	DB *gorm.DB
	Repository repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceBio{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceBio) GetByID(c echo.Context, id uint) (biostar.BioStarResponseDTO, error) {
	var biostarRes biostar.BioStarResponseDTO

	result, err := s.Repository.GetByID(c, s.DB, id)
	if err != nil {
		return biostarRes, err
	}

	biostarRes = result.ToResponse()

	return biostarRes, nil
}

func (s *serviceBio) GetAll(c echo.Context) ([]biostar.BioStarResponseDTO, error) {
	var biostarRes []biostar.BioStarResponseDTO

	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return biostarRes, err
	}

	for _, biostar := range result {
		biostarRes = append(biostarRes, biostar.ToResponse())
	}

	return biostarRes, nil
}