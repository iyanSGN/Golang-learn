package warga

import "time"

type WargaRequestDTO struct {
	ID          uint   	`json:"id"`
	IDKecamatan *uint   `json:"id_kecamatan" validate:"required"`
	IDKabupaten *uint	`json:"id_kabupaten" validate:"required"`
	IDProvinsi	*uint	`json:"id_provinsi" validate:"required"`
	Nama        string 	`json:"nama" validate:"required"`
	NoKtp		string	`json:"no_ktp" validate:"required"`
}

type WargaResponseDTO struct {
	ID          	uint   		`json:"id"`
	Nama        	string 		`json:"name"`
	NoKtp       	string 		`json:"no_ktp"`
	IDKecamatan 	*uint   	`json:"id_kecamatan"`
	Kecamatan_nama	string		`json:"kecamatan_nama"`
	IDKabupaten 	*uint		`json:"id_kabupaten"`
	Kabupaten_nama	string		`json:"Kabupaten_nama"`
	IDProvinsi		*uint		`json:"id_provinsi"`
	Provinsi_nama	string		`json:"Provinsi_nama"`
	CreatedAt   	time.Time	`json:"created_at"`

}