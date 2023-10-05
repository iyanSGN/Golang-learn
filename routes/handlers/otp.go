package handlers

import (
	"fmt"
	"net/http"
	"rearrange/middleware"
	"rearrange/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GenerateOTP(c echo.Context) error {
	var admin	models.MOTP
	if err := c.Bind(&admin);
	err != nil {
		return err
	}


	repo := middleware.NewOtpRepository()

	isVerified, verifyErr := repo.IsUserVerified(admin.CreatedBy)
	if verifyErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to check admin verification status",
		})
	}

	if isVerified {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User is already verified",
		})
	}

	email, err := repo.GetAdminEmailByID(admin.CreatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed check admin verification status",
		})
	}
	
	createOtp := repo.GenerateOTP()
	createErr := repo.CreateOtp(admin.CreatedBy, createOtp)

	if createErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error" : "failed to generate and store otp, try to resend otp",
			"status_code" : http.StatusInternalServerError,
		})
	}

	createErr = middleware.Otp(email, strconv.Itoa(int(createOtp)))
	if createErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error" : "Failed to send otp email",
		})
	}
	

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Otp sent Success",
		"status_code" : http.StatusOK,
	})

}

func ResendOTP(c echo.Context) error {
	var admin models.MOTP
	if err := c.Bind(&admin);

	err != nil {
		return err
	}

	repo := middleware.NewOtpRepository()

	email, err := repo.GetAdminEmailByID(admin.CreatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error" : "failed to retrieve admin email",
		})
	}

	isverified, err :=  repo.IsUserVerified(admin.CreatedBy)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to check user verification status",
		})
	}

	if isverified {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "User is already verified",
			"status_code": http.StatusOK,
		})
	}

	UpdatedOtp :=  repo.GenerateOTP()
	createErr :=  repo.UpdateorCreateOtp(admin.CreatedBy, UpdatedOtp)
	if createErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":       "Failed to resend OTP",
			"status_code": http.StatusInternalServerError,
		})
	}

	createErr = middleware.Otp(email, strconv.Itoa(int(UpdatedOtp)))
	if createErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to send OTP email",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message" : "otp resent successfully",
		"status_code" : http.StatusOK,
	})

}

func VerifyOtp(c echo.Context) error {
	var admin models.MOTP
	if err := c.Bind(&admin); err != nil {
		return err
	}

	otpRepo := middleware.NewOtpVerify()

	otpVerification, err := otpRepo.VerifyOtp(admin.CreatedBy, admin.OtpCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Error verifying OTP",
		})
	}

	if otpVerification {
		err := middleware.UpdateUserIsActive(admin.CreatedBy)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "Error updating user activation",
			})
		}

		middleware.UpdateOtp(admin.CreatedBy)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message" : "otp verification is successfull. admin is active",
			"status_code" : http.StatusOK,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"Error" : "invalid otp",
	})
}

func ScheduleDeleteExpiredOtp() {
	for {
		repo := middleware.NewOtpRepository()
		err := repo.DeleteExpiredOtp()
		if err != nil {
			fmt.Println("Error deleting expired otp:", err)
		}
		time.Sleep(5 * time.Minute)
	}
}