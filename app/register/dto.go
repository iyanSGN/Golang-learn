package register

import "time"

type AdminRequestDTO struct {
	ID       uint   `json:"id"`
	RoleID	 *uint	`json:"role_id" validate:"required"`
	Nama     string `json:"nama" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminResponseDTO struct {
	ID        uint   		`json:"id"`
	Email     string 		`json:"email"`
	Password  string 		`json:"password"`
	CreatedAt time.Time		`json:"created_at"`	
}