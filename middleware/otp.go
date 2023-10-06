package middleware

import (
	"math/rand"
	"rearrange/models"
	"rearrange/package/database"
	"time"

	"gorm.io/gorm"
)

type OtpRepository struct{}

func NewOtpRepository() *OtpRepository {
	return &OtpRepository{}
}

func (repo *OtpRepository) GenerateOTP() int32 {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	min := 1000000
	max := 9999999
	return int32(rng.Intn(max-min + 1) + min)
}

func (repo *OtpRepository)CreateOtp(adminID int32, otp int32) error {
	db := database.GetDB()

	newOtp :=  models.MOTP {
		OtpCode: otp,
		CreatedBy: adminID,
		Status: "unverif",
	}

	return db.Create(&newOtp).Error
}

func (repo *OtpRepository) UpdateorCreateOtp(adminID int32, otp int32) error {
	db := database.GetDB()

	return db.Transaction(func(tx *gorm.DB) error  {
		var exisingOtp models.MOTP
		if err :=  tx.Where("created_by = ?", adminID).First(&exisingOtp).Error;
		err != nil {
			if err == gorm.ErrRecordNotFound {
				return	repo.CreateOtp(adminID, otp)
			}

			return err
		}

		return tx.Model(&exisingOtp).Update("otp_code", otp).Error
	})
}

func (repo *OtpRepository)IsUserVerified(adminID int32) (bool, error) {
	db := database.GetDB()

	var admin models.MRegister
	if err := db.Select("isactive").Where("id = ?", adminID).First(&admin).Error;

	err != nil { 
		return false, err
	}

	return admin.Isactive == 1, nil
}

func (repo *OtpRepository) DeleteExpiredOtp()error {
	db :=  database.GetDB()

	expirationTime := time.Now().Add(-5 * time.Minute)
	err := db.Where("createdat < ?", expirationTime).Delete(&models.MOTP{}).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	return err
}

func (repo *OtpRepository) GetAdminEmailByID(adminID int32) (string, error) {
	db := database.GetDB()

	var admin models.MRegister
	if err := db.Select("email").Where("id= ?", adminID).First(&admin).Error;

	err != nil {
		return "", err
	}

	return admin.Email, nil
}