package provinsi

import(
	"time"
)

type ProvinsiRequestDTO struct {
	Nama			string		`json:"nama" validate:"required"`
	CreatedBy		uint		`json:"createdby"`
	UpdatedBy		uint		`json:"updatedby"`

}

type ProvinsiResponseDTO struct {
	ID			uint		`json:"id"`
	Nama		string		`json:"nama"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
	CreatedBy	uint		`json:"created_by"`
}