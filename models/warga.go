package models

import (
	"rearrange/app/warga"
	"time"
)

type MWarga struct {
	ID        	uint   		`json:"id" gorm:"primary_key"`
	IDKecamatan	*uint		`json:"id_kecamatan"`
	IDKabupaten	*uint		`json:"id_kabupaten"`
	IDProvinsi	*uint		`json:"id_provinsi"`
	Nama     	string 		`json:"nama" gorm:"varchar(100)"`
	NoKtp     	string 		`json:"no_ktp" gorm:"varchar(100);unique"`
	IsActive  	bool   		`json:"is_active" gorm:"default:true"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`

	Kecamatan	*MKecamatan	`json:"kecamtan" gorm:"foreignKey:IDKecamatan"`
	Kabupaten	*MKabKota	`json:"kabupaten" gorm:"foreignKey:IDKabupaten"`
	Provinsi	*MProvinsi	`json:"provinsi" gorm:"foreignKey:IDProvinsi"`


}


func (mk *MWarga) ToResponse() warga.WargaResponseDTO {
	return warga.WargaResponseDTO{
		ID: mk.ID,
		Nama: mk.Nama,
		NoKtp: mk.NoKtp,
		IDKecamatan: mk.IDKecamatan,
		IDKabupaten: mk.IDKabupaten,
		IDProvinsi: mk.IDProvinsi,
		CreatedAt: mk.CreatedAt,
	}
}