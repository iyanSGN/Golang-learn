package middleware

import (
	"fmt"
	"rearrange/models"
	"rearrange/package/database"
	"time"

	"gorm.io/gorm"
)

type OtpVerify struct {
}

func NewOtpVerify() *OtpVerify {
	return &OtpVerify{}
}

func (verify *OtpVerify) VerifyOtp(adminID int32, otpCode int32) (bool, error) {
	db := database.GetDB()

	var exisitngOtp models.MOTP
	if err := db.Where("created_by = ?", adminID).First(&exisitngOtp).Error;
	err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("otp record not found for admin:",adminID)
			return false, nil
		}
		fmt.Println("error fetching otp record:", err)
		return false, nil
	}

	if exisitngOtp.OtpCode == otpCode {
		expirationTime :=  time.Now().Add(-10 *time.Minute)
		if exisitngOtp.CreatedAt.Before(expirationTime) {
			fmt.Println("OTP expired for admin:", adminID)
			return false, nil
		}
		fmt.Println("OTP verified for admin:", adminID)
		return true, nil
	}

	fmt.Println("otp does not match for admin:", adminID)
	return false, nil
}

func UpdateUserIsActive(adminID int32) error {
	db := database.GetDB()

	result := db.Model(&models.MRegister{}).
				Where("id = ?", adminID).
				Updates(map[string]interface{}{
					"isactive" : 1,
					"updated_at" : time.Now(),
				})

			if result.Error != nil {
				return result.Error
			}

			return nil
}

func UpdateOtp(adminID int32) error {
	db := database.GetDB()

	result := db.Model(&models.MOTP{}).
				Where("created_by = ?", adminID).
				Update("status", "Verified")

		if result.Error != nil {
			return result.Error
		}

		return nil
}