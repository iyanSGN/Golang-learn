package repository

import (
	"fmt"
	dto "rearrange/app/user"
	"rearrange/models"
	"rearrange/package/database"
	"time"

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

func CreateUser(newUser dto.PostUser, imageProfile []byte, formatImage string) (models.MRegister, error) {
	db := database.GetDB()

	user := models.MRegister {
		ID: newUser.ID,
		CreatedBy: newUser.CreatedBy,
		CreatedAt: newUser.CreatedAt,
		UpdatedBy: newUser.UpdatedBy,
		UpdatedAt: newUser.CreatedAt,
		Nama: newUser.Nama,
		Phone: newUser.Phone,
		Email: newUser.Email,
		Password: newUser.Password,
		ImageProfile: imageProfile,
		FormatProfile: formatImage,

	}

	result := db.Create(&user)
	if result.Error != nil {
		return user, fmt.Errorf("error creating user: %w", result.Error)
	}

	return user, nil
}

func UpdateAdmin(id uint, adminID models.MRegister ) error {
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

func DefaultPicture(id int, imageProfile []byte, image dto.ImageUser, formatImage string) error {
	db := database.GetDB()

	var user models.MRegister

	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return fmt.Errorf("failed to find : %w", err)
	}

	fmt.Println(imageProfile)

	user.FormatProfile = formatImage
	user.ImageProfile = imageProfile
	user.UpdatedBy = image.UpdatedBy
	user.UpdatedAt = time.Now()

	updateResult := db.Save(&user)
	if updateResult.Error != nil {
		return fmt.Errorf("failed upload image: %w", updateResult.Error)
	}
	return nil
}