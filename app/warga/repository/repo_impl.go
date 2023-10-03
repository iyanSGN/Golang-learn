package repository

import (
	"fmt"
	"rearrange/app/warga"
	"rearrange/models"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepo() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.MWarga, error) {
	var Warga []models.MWarga

	if err := DB.Preload("Kecamatan.KabKota.Provinsi").Find(&Warga).Error;
	err != nil {
		return nil, err
	}

	return Warga, nil
}

func (r *repositoryImpl)GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.MWarga, error) {
	var warga	models.MWarga

	if err :=  DB.Where("id = ?", ID).Preload("Kecamatan.KabKota.Provinsi").First(&warga).Error;
	err != nil {
		return warga, err
	}

	return warga, nil
}

func CreateWarga(request models.MWarga) (models.MWarga, error) {
	db := database.GetDB()

	warga := warga.WargaRequestDTO{
		ID: request.ID,
	}

	if request.IDKecamatan != nil {
		kecamatanID := uint(*request.IDKecamatan)
		warga.IDKecamatan = &kecamatanID
	}

	if request.IDKabupaten != nil {
		kabupatenID := uint(*request.IDKabupaten)
		warga.IDKabupaten = &kabupatenID
	}

	if request.IDProvinsi != nil {
		provinsiID := uint(*request.IDProvinsi)
		warga.IDProvinsi = &provinsiID
	}

	result := db.Create(&request)
	if result.Error != nil {
		return request, fmt.Errorf("error creating warga: %w", result.Error)
	}

	return request, nil
}

func UpdateWarga(id int, request warga.WargaRequestDTO) error {
	db := database.GetDB()

	var warga models.MWarga

	result := db.First(&warga, id)
	if result.Error != nil {
		return fmt.Errorf("error updating warga: %w", result.Error)
	}

	if request.Nama != "" {
		warga.Nama = request.Nama
	}

	if request.NoKtp != "" {
		warga.NoKtp = request.NoKtp
	}

	if request.IDKecamatan != nil{
		kecamatanID := uint(*request.IDKecamatan)
		warga.IDKecamatan = &kecamatanID
	}

	if request.IDKabupaten != nil {
		kabupatenID := uint(*request.IDKabupaten)
		warga.IDKabupaten = &kabupatenID
	}

	if request.IDProvinsi != nil {
		provinsiID := uint(*request.IDProvinsi)
		warga.IDProvinsi = &provinsiID
	}

	updatedWarga := db.Save(&warga)
	if updatedWarga.Error != nil {
		return fmt.Errorf("error updatig warga: %w", updatedWarga.Error)
	}

	return nil
}

func DeleteWarga(id uint) error {
	db := database.GetDB()

	var admin models.MWarga

	result := db.First(&admin, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving warga: %w", result.Error)
	}

	deleteWarga := db.Delete(&admin)
	if deleteWarga.Error != nil {
		return fmt.Errorf("error deleting user: %w", deleteWarga.Error)
	}

	return nil 
}