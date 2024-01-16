package models

import (
	"rearrange/app/kecamatan"
	"time"
)

type MKecamatan struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	IDKabKota *uint     `json:"id_kab_kota" gorm:"not null"`
	Nama      string    `json:"nama" gorm:"type:varchar(100);not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	KabKota MKabKota `json:"kab_kota" gorm:"foreignKey:IDKabKota"`
}

func (mk *MKecamatan) ToResponse() kecamatan.KecamatanResponseDTO {
	return kecamatan.KecamatanResponseDTO{
		ID:           mk.ID,
		KabKota_id:   mk.IDKabKota,
		KabKota_nama: mk.KabKota.Nama,
		ProvinceID:   mk.KabKota.Provinsi.ID,
		ProvinceName: mk.KabKota.Provinsi.Nama,
		Nama:         mk.Nama,
		CreatedAt:    mk.CreatedAt.UTC(),
	}
}