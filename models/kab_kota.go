package models

import (
	"rearrange/app/kabupaten"
	"time"
)

type MKabKota struct {
	ID        		uint  			`json:"id" gorm:"primary_key"`
	IDProvinsi      *uint  			`json:"id_Provinsi" gorm:"not null"`
	Nama	   		string 			`json:"nama" gorm:"type:varchar(100);not null"`
	IsActive  		bool   			`json:"is_active" gorm:"default:true"`
	CreatedBy 		uint   			`json:"created_by"`
	UpdatedBy 		uint   			`json:"updated_by"`
	CreatedAt 		time.Time		`json:"created_at"`
	UpdatedAt 		time.Time 		`json:"updated_at"`

	Provinsi		*MProvinsi		`json:"provinsi" gorm:"foreignKey:IDProvinsi"`
}

func (mk *MKabKota) ToResponse() kabupaten.KabKotaResponseDTO {
	return kabupaten.KabKotaResponseDTO{
		ID: mk.ID,
		Provinsi_id: mk.IDProvinsi,
		Provinsi_nama: mk.Provinsi.Nama,
		Nama: mk.Nama,
		CreatedAt: mk.CreatedAt,
	}
}


