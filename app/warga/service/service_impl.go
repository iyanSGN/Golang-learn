package service

import (
	service "rearrange/app/warga"
	"rearrange/app/warga/repository"
	"rearrange/package/response"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB 			*gorm.DB
	Repository 	repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl) GetAll(c echo.Context) ([]service.WargaResponseDTO,error) {
	var wargaRes []service.WargaResponseDTO

	result, err :=  s.Repository.GetAll(c, s.DB)
	if err != nil {
		return wargaRes, response.BuildError(response.ErrServerError, err)
	}

	for _, warga := range result {
		wargaRes = append(wargaRes, warga.ToResponse())
	}

	return wargaRes, nil

}

func (s *serviceImpl) GetByID(c echo.Context, ID uint) (service.WargaResponseDTO, error) {
	var admin service.WargaResponseDTO

	result, err := s.Repository.GetByID(c, s.DB, ID)
	if err != nil {
		return admin, response.BuildError(response.ErrServerError,err)
	}

	admin = result.ToResponse()

	return admin, nil
	
}