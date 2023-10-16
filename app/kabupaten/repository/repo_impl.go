package repository

import (
	"fmt"
	"rearrange/app/kabupaten"
	"rearrange/models"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.MKabKota,error) {
	var kabupaten []models.MKabKota

	if err := DB.Preload("Provinsi").Find(&kabupaten).Error; err != nil {
		return nil, err
	}

	return kabupaten, nil
}

func (r *repositoryImpl) GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.MKabKota, error) {
	var kabupaten models.MKabKota

	if err := DB.Where("id = ?", ID).Preload("Provinsi").First(&kabupaten).Error;
	err != nil {
		return kabupaten, err
	}

	return kabupaten, nil
}

func CreateKabupaten(request kabupaten.KabKotaRequestDTO) (kabupaten.KabKotaRequestDTO, error) {
	db := database.GetDB()

	kabupaten := models.MKabKota{
		IDProvinsi: request.IDProvinsi,
		Nama: request.Nama,
		CreatedBy: request.CreatedBy,
		UpdatedBy: request.CreatedBy,
	}

	if request.IDProvinsi != nil {
		ProvinsiID := uint(*request.IDProvinsi)
		kabupaten.IDProvinsi = &ProvinsiID
	}

	result := db.Create(&kabupaten)
	if result.Error != nil {
		return request, fmt.Errorf("error creating Kabupaten: %w", result.Error)
	}

	return request, nil
}

func UpdateKabupaten(id int, request kabupaten.KabKotaRequestDTO) error {
	db := database.GetDB()

	var kab  models.MKabKota
	
	result := db.First(&kab, id)
	if result.Error != nil {
		return fmt.Errorf("error updating kabupaten: %w", result.Error)
	}

	if request.Nama != "" {
		kab.Nama = request.Nama
	}

	if request.IDProvinsi != nil {
		ProvinsiID := uint(*request.IDProvinsi)
		kab.IDProvinsi = &ProvinsiID
	}

	UpdatedKabupaten := db.Save(&kab)
	if UpdatedKabupaten.Error != nil {
		return fmt.Errorf("error updating Kabupaten: %w", UpdatedKabupaten.Error)
	}

	return nil
}

func DeleteKabupaten(id uint) error {
	db := database.GetDB()

	var kab models.MKabKota

	result := db.First(&kab, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving kabupaten: %w", result.Error)
	}

	deleteKab := db.Delete(&kab)
	if deleteKab.Error != nil {
		return fmt.Errorf("error deleting user: %w", deleteKab.Error)
	}

	return nil

}