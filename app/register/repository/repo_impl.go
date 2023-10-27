package repository

import (
	"fmt"
	"rearrange/models"
	"rearrange/package/database"
)

func CreateAdmin(admin models.MRegister) (models.MRegister, error) {
	db := database.GetDB()

	result := db.Create(&admin)
	if result.Error != nil {
		return admin, fmt.Errorf("error create admin: %w", result.Error)
	}
	return admin, nil
}
