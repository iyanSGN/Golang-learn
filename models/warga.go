package models

import (
	"rearrange/app/warga"
	"time"
)

type MWarga struct {
	ID        	uint   		`json:"id" gorm:"primary_key"`
	IDKecamatan	*uint		`json:"id_kecamatan"`
	Nama     	string 		`json:"nama" gorm:"varchar(100)"`
	NoKtp     	string 		`json:"no_ktp" gorm:"varchar(100);unique"`
	IsActive  	bool   		`json:"is_active" gorm:"default:true"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"updated_at"`

	Kecamatan	*MKecamatan	`json:"kecamtan" gorm:"foreignKey:IDKecamatan"`


}


func (mk *MWarga) ToResponse() warga.WargaResponseDTO {
	return warga.WargaResponseDTO{
		ID: mk.ID,
		Nama: mk.Nama,
		NoKtp: mk.NoKtp,
		IDKecamatan: mk.IDKecamatan,
		Kecamatan_nama: mk.Kecamatan.Nama,
		IDKabupaten: mk.Kecamatan.IDKabKota,
		Kabupaten_nama: mk.Kecamatan.KabKota.Nama,
		IDProvinsi: mk.Kecamatan.KabKota.IDProvinsi,
		Provinsi_nama: mk.Kecamatan.KabKota.Provinsi.Nama,
		CreatedAt: mk.CreatedAt,
	}
}