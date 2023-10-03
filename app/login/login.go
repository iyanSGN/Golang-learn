package login

import (
	"fmt"
	"rearrange/models"
	"rearrange/package/database"

)

func GetAdminByEmail(email string) (*models.MRegister, error) {
	db :=  database.GetDB()
	admin := &models.MRegister{}

	result := db.Where("email = $1", email).First(admin)
	if result.Error != nil {
		return nil, fmt.Errorf("error fetching akun : %w", result.Error)
	}

	return admin, nil
}