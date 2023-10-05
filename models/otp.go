package models

import "time"

type MOTP struct {
	Id        	int32 		`gorm:"primarykey" json:"id"`
	CreatedAt 	time.Time	`gorm:"default:current_timestamp" json:"created_at"`
	CreatedBy	int32		`gorm:"unique" json:"createdby"`
	OtpCode		int32		`json:"otp_code"`
	Status		string		`json:"status"`	
}

