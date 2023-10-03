package models

import (
	"rearrange/app/register"
	"time"
)

type MRegister struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Nama      string    `json:"nama"`
	Phone     string    `json:"phone"`
	Email     string    `gorm:"unique;type:varchar(255)" json:"email"`
	Password  string    `json:"password"`
	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (mk *MRegister) ToResponse() register.AdminResponseDTO{
	return register.AdminResponseDTO{
		ID: mk.ID,
		Email: mk.Email,
		Password: mk.Password,
		CreatedAt: mk.CreatedAt,
	}
}