package provinsi

import(
	"time"
)

type ProvinsiRequestDTO struct {
	ID 		uint 		`json:"id"`
	Nama	string		`json:"nama" validate:"required"`

}

type ProvinsiResponseDTO struct {
	ID			uint		`json:"id"`
	Nama		string		`json:"nama"`
	CreatedAt 	time.Time	`json:"created_at"`
}