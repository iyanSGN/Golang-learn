package user


import "time"

type GetUser struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	FormatProfile  string `json:"format_profile"`
}

type PostUser struct {
	ID       	uint   		`form:"id"`
	CreatedBy	uint     	`form:"createdby"`
	CreatedAt   time.Time 	`form:"createdat"`
	UpdatedBy   uint     	`form:"updatedby"`
	UpdatedAt   time.Time 	`form:"updatedat"`
	Nama     	string 		`form:"nama"`
	Phone    	string 		`form:"phone"`
	Email    	string 		`form:"email"`
	Password 	string 		`form:"password"`
	ImageProfile  []byte 	`form:"image_profile"`
	FormatProfile string    `form:"format_profile"`
}

type PutUser struct {
	Nama     string `json:"nama" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	ImageProfile  []byte    `form:"image_profile"`
}

type UserResponseDTO struct {
	ID        uint   		`json:"id"`
	Email     string 		`json:"email"`
	Password  string 		`json:"password"`
	CreatedAt time.Time		`json:"created_at"`	
}

type ImageUser struct {
	UpdatedBy    uint       `json:"updatedby"`
	UpdatedAt    time.Time `json:"updatedat"`
	ImageProfile string    `json:"image_profile"`
	FormatImage  string    `json:"format_image"`
}