package service

import (
	"rearrange/app/provinsi"
	"rearrange/app/provinsi/repository"
	"rearrange/package/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB 				*gorm.DB
	Repository 		repository.Repository
}

func NewService (DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) GetAll(c echo.Context) ([]provinsi.ProvinsiResponseDTO,error) {
	var provinsiRes []provinsi.ProvinsiResponseDTO

	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return provinsiRes, response.BuildError(response.ErrServerError, err)
	}

	for _, provinsi := range result {
		provinsiRes = append(provinsiRes, provinsi.ToResponse())
	}

	return provinsiRes, nil

}

func (s *serviceImpl) GetByID(c echo.Context, ID uint) (provinsi.ProvinsiResponseDTO, error) {
	var provinsi provinsi.ProvinsiResponseDTO

	result, err := s.Repository.GetByID(c, s.DB, ID)
	if err != nil {
		return provinsi, response.BuildError(response.ErrNotFound, err)
	}

	provinsi = result.ToResponse()

	return provinsi, nil
}
