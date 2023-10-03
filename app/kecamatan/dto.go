package kecamatan

import "time"

type KecamatanRequestDTO struct {
	ID        uint   `json:"id"`
	IDKabKota *uint  `json:"id_kab_kota" validate:"required"`
	Nama      string `json:"nama" validate:"required"`
}

type KecamatanResponseDTO struct {
	ID           uint      `json:"id"`
	Nama         string    `json:"nama"`
	CreatedAt    time.Time `json:"created_at"`
	KabKota_id   *uint     `json:"kab_kota"`
	KabKota_nama string    `json:"kabupaten_nama"`
	ProvinceID   uint      `json:"province_id"`
	ProvinceName string    `json:"province_name"`
}
