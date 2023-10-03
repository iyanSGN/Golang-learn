package service

import (
	"rearrange/app/register"
	"rearrange/app/register/repository"
	"rearrange/package/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type serviceImpl struct {
	DB *gorm.DB
	Repository repository.Repository
}

func NewService(DB *gorm.DB, Repository repository.Repository) Service {
	return &serviceImpl{
		DB: DB,
		Repository: Repository,
	}
}

func (s *serviceImpl)GetAll(c echo.Context) ([]register.AdminResponseDTO, error) {
	var adminRes []register.AdminResponseDTO

	result, err := s.Repository.GetAll(c, s.DB)
	if err != nil {
		return adminRes, response.BuildError(response.ErrServerError, err)
	}

	for _, admin := range result {
		adminRes = append(adminRes, admin.ToResponse())
	}

	return adminRes, nil
}


func (s *serviceImpl)GetByID(c echo.Context, id uint) (register.AdminResponseDTO, error) {
	var adminRes register.AdminResponseDTO

	result, err := s.Repository.GetByID(c, s.DB, id)
	if err != nil {
		return adminRes, response.BuildError(response.ErrNotFound, err)
	}

	adminRes = result.ToResponse()

	return adminRes, nil
}