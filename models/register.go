package models

import (
	"rearrange/app/user"
	"time"
)

type MRegister struct {
	ID        		uint      		`gorm:"primarykey" json:"id"`
	Isactive  		int32     		`json:"isactive"`
	Nama      		string    		`json:"nama"`
	Phone     		string    		`json:"phone"`
	Email     		string    		`gorm:"type:varchar(255)" json:"email"`
	Password  		string    		`json:"password"`
	ImageProfile 	[]byte    		`gorm:"type:bytea" json:"image_profile"`
	FormatProfile 	string    		`gorm:"type:varchar(255)" json:"format_profile"`
	CreatedBy 		uint      		`json:"createdby"`
	UpdatedBy 		uint      		`json:"updatedby"`
	CreatedAt 		time.Time 		`json:"createdat"`
	UpdatedAt 		time.Time 		`json:"updatedat"`
	Provinsi1		MProvinsi		`gorm:"foreignKey:CreatedBy"`
	Provinsi2		MProvinsi		`gorm:"foreignKey:UpdatedBy"`
	Kabupaten1		MKabKota		`gorm:"foreignKey:CreatedBy"`
	Kabupaten2		MKabKota		`gorm:"foreignKey:UpdatedBy"`
	Kecamatan1		MKecamatan		`gorm:"foreignKey:CreatedBy"`
	Kecamatan2		MKecamatan		`gorm:"foreignKey:UpdatedBy"`
	

		
}

func (mk *MRegister) ToResponse() user.UserResponseDTO{
	return user.UserResponseDTO{
		ID: mk.ID,
		Email: mk.Email,
		Password: mk.Password,
		CreatedAt: mk.CreatedAt,
	}
}