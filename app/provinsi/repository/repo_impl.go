package repository

import (
	"fmt"
	"rearrange/app/provinsi"
	"rearrange/models"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.MProvinsi,error) {
	var provinsi []models.MProvinsi

	if err := DB.Find(&provinsi).Error; err != nil {
		return nil, err
	}

	return provinsi, nil
}

func (r *repositoryImpl) GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.MProvinsi, error) {
	var provinsi models.MProvinsi

	if err := DB.Where("id = ?", ID).First(&provinsi).Error; 
	err != nil {
		return provinsi, nil
	}

	return provinsi, nil
}

func CreateProvinsi(request provinsi.ProvinsiRequestDTO ) (provinsi.ProvinsiRequestDTO, error) {
	db := database.GetDB()


	postProvinsi := models.MProvinsi{
		Nama: request.Nama,
		CreatedBy: request.CreatedBy,
		UpdatedBy: request.UpdatedBy,
	}

	result := db.Create(&postProvinsi)
	if result.Error != nil {
		return request, fmt.Errorf("error creating provinsi: %w", result.Error)
	}

	return request, nil
}

func UpdateProvinsi(id int, provID models.MProvinsi) error {
	db := database.GetDB()

	var provinsi models.MProvinsi
	result := db.First(&provinsi, id)

	if result.Error != nil {
		return fmt.Errorf("error: %w", result.Error)
	}

	if provID.Nama != "" {
		provinsi.Nama = provID.Nama
	}

	updatedProvinsi := db.Save(&provinsi)
	if updatedProvinsi.Error != nil {
		return fmt.Errorf("error saving updates: %w", updatedProvinsi.Error)
	}

	return nil
}


func DeleteUser(id int) error {
	db := database.GetDB()

	var provinsi models.MProvinsi
	result := db.Delete(&provinsi, id)
	if result.Error != nil {
		return fmt.Errorf("error : %w", result.Error)
	}

	return nil
}