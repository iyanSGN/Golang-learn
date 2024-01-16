package helpers

import (
	dto "rearrange/app/user"
	"rearrange/models"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(user *models.MRegister) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}

func HashedPassword(Post *dto.PostUser) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Post.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	Post.Password = string(hashedPassword)
	return nil
}