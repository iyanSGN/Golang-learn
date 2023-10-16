package repository

import (
	"fmt"
	"rearrange/app/kecamatan"
	"rearrange/models"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repositoryImpl struct{}

func NewRepository() Repository {
	return &repositoryImpl{}
}

func (r *repositoryImpl) GetAll(c echo.Context, DB *gorm.DB) ([]models.MKecamatan, error) {
	var kecamatan []models.MKecamatan

	if err := DB.Preload("KabKota.Provinsi").Find(&kecamatan).Error; 
	err != nil {
		return nil, err
	}

	return kecamatan, nil
}

func (r *repositoryImpl) GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.MKecamatan, error) {
	var kecamatan models.MKecamatan

	if err := DB.Where("ID = ?", ID).Preload("KabKota.Provinsi").First(&kecamatan).Error;
	err != nil {
		return kecamatan, nil
	}

	return kecamatan, nil
}

func CreateKecamatan(request kecamatan.KecamatanRequestDTO) (kecamatan.KecamatanRequestDTO, error) {
	db := database.GetDB()

	kecamatan := models.MKecamatan{
		IDKabKota: request.IDKabKota,
		Nama: request.Nama,
		CreatedBy: request.CreatedBy,
		UpdatedBy: request.UpdatedBy,
	}

	if request.IDKabKota != nil {
		kecID := uint(*request.IDKabKota)
		kecamatan.IDKabKota = &kecID
	}

	result := db.Create(&kecamatan)
	if result.Error != nil {
		return request, fmt.Errorf("error creating Kabupaten: %w", result.Error)
	}

	return request, nil


}

func UpdateKecamatan(id int, request kecamatan.KecamatanRequestDTO) error {
	db := database.GetDB()

	var kec models.MKecamatan

	result := db.First(&kec, id)
	if result.Error != nil {
		return fmt.Errorf("error updating kecamatan: %w", result.Error)
	}

	if request.Nama != "" {
		kec.Nama = request.Nama
	}

	if request.IDKabKota != nil {
		KabupatenID := uint(*request.IDKabKota)
		kec.IDKabKota = &KabupatenID
	}

	updatedKecamatan := db.Save(&kec)
	if updatedKecamatan.Error != nil {
		return fmt.Errorf("error updating Kecamatan: %w", updatedKecamatan.Error)
	}

	return nil
}

func DeleteKecamatan(id int) error {
	db := database.GetDB()

	var kec models.MKecamatan

	result := db.First(&kec, id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving Kecamatan: %w", result.Error)
	}

	deleteKec := db.Delete(&kec)
	if deleteKec.Error != nil {
		return fmt.Errorf("error deleting kecamatan: %w", deleteKec.Error)
	}

	return nil

}
