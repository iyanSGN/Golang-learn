package kabupaten

import "time"

type KabKotaRequestDTO struct {
	ID         	uint `json:"id"`
	IDProvinsi 	*uint `json:"id_provinsi" validate:"required"`
	Nama       	string `json:"nama" validate:"required"`
}

type KabKotaResponseDTO struct {
	ID         		uint 			`json:"id"`
	Nama      		string 			`json:"nama"`
	CreatedAt  		time.Time		`json:"createdAt"`
	Provinsi_id   	*uint			`json:"provinsi_id"`
	Provinsi_nama	string			`json:"provinsi_nama"`
}

