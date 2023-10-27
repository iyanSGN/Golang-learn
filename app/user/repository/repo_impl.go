package repository

import (
	"fmt"
	"rearrange/models"
	"rearrange/package/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type repoImpl struct {
}

func NewRepo() Repository {
	return &repoImpl{}
}

func (r *repoImpl) GetAll(c echo.Context, DB *gorm.DB ) ([]models.MRegister, error) {
	var	admin []models.MRegister

	if err := DB.Find(&admin).Error;
	err != nil {
		return nil, err
	}

	return admin, nil
}

func (r *repoImpl) GetByID(c echo.Context, DB *gorm.DB, ID uint) (models.MRegister, error) {
	var admin	models.MRegister

	if err := DB.Where("id= ?", ID).First(&admin).Error;
	err != nil {
		return admin, nil
	}

	return admin, nil
}

func UpdateAdmin(id int, adminID models.MRegister ) error {
	db := database.GetDB()

	var admin models.MRegister
	result := db.First(&admin, id)

	if result.Error != nil {
		return fmt.Errorf("error : %w", result.Error)
	}

	if adminID.Nama != "" {
		admin.Nama = adminID.Nama
	}

	if adminID.Phone != "" {
		admin.Phone = adminID.Phone
	}

	if adminID.Email != "" {
		admin.Email = adminID.Email
	}

	if adminID.Password != "" {
		admin.Password = adminID.Password
	}

	if adminID.RoleID != nil {
		roleID := uint(*adminID.RoleID)
		admin.RoleID = &roleID
	}

	updatedAdmin := db.Save(&admin)
	if updatedAdmin.Error != nil {
		return fmt.Errorf("error saving updates: %w", updatedAdmin.Error)
	}

	return nil

	


}

func DeleteAdmin(id int) error {
	db := database.GetDB()

	var adminID models.MRegister
	result :=  db.First(&adminID,id)
	if result.Error != nil {
		return fmt.Errorf("error retrieving selected user: %w", result.Error)
	}

	deleteAdmin := db.Delete(&adminID)
	if deleteAdmin.Error != nil {
		return fmt.Errorf("error deleting selected user: %w", deleteAdmin.Error)
	}

	return nil
}