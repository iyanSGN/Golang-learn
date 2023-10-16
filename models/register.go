package models

import (
	"rearrange/app/register"
	"time"
)

type MRegister struct {
	ID        		uint      		`json:"id" gorm:"primary_key"`
	RoleID	  		*uint			`json:"role_id"`
	Isactive  		int32     		`json:"isactive"`
	Nama      		string    		`json:"nama"`
	Phone     		string    		`json:"phone"`
	Email     		string    		`gorm:"unique;type:varchar(255)" json:"email"`
	Password  		string    		`json:"password"`
	CreatedBy 		uint      		`json:"createdby"`
	UpdatedBy 		uint      		`json:"updatedby"`
	CreatedAt 		time.Time 		`json:"createdat"`
	UpdatedAt 		time.Time 		`json:"updatedat"`
	Provinsi1		MProvinsi		`gorm:"foreignKey:CreatedBy"`
	Provinsi2		MProvinsi		`gorm:"foreignKey:UpdatedBy"`
}

func (mk *MRegister) ToResponse() register.AdminResponseDTO{
	return register.AdminResponseDTO{
		ID: mk.ID,
		Email: mk.Email,
		Password: mk.Password,
		CreatedAt: mk.CreatedAt,
	}
}