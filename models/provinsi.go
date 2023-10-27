package models

import (
	"rearrange/app/provinsi"
	"time"
)

type MProvinsi struct {
	ID 			uint 		`json:"id" gorm:"primary_key"`
	Nama		string		`json:"nama" gorm:"type:varchar(100);not null"`
	CreatedBy 	uint     	`json:"created_by"`
	UpdatedBy 	uint      	`json:"updated_by"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

func (mk *MProvinsi) ToResponse() provinsi.ProvinsiResponseDTO {
	return provinsi.ProvinsiResponseDTO{
		ID: mk.ID,
		Nama: mk.Nama,
		CreatedAt: mk.CreatedAt,
		UpdatedAt: mk.UpdatedAt,
		CreatedBy: mk.CreatedBy,

	}
}