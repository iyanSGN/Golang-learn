package models

import (
	"rearrange/app/biostar"
	"time"
)

type MBioStar struct {
	UserId     			string   		`json:"userid"`
	Name          		string 			`json:"name"`
	UserGroupId			uint   			`json:"usergroup_id" gorm:"primary_key"`
	Disabled     		bool   			`json:"disabled"`
	StartDateTime 		time.Time		`json:"startdt"`
	ExpiryDatetime 		time.Time		`json:"expirydt"`
	Email				string			`json:"email"`
	Department			string			`json:"department" gorm:"type:varchar(255)"`
	Title				string			`json:"title"`
	Photo				[]byte			`json:"photo"`
	Phone				string			`json:"phone"`
	Permission			uint			`json:"permission"`
	AccessGroup     	uint			`json:"accessgroup"`
	LoginID				string			`json:"loginid"`
	Password			string			`json:"password"`
	UserIp				string			`json:"userip"`
	Pin					uint			`json:"pin"`
}

func (mk *MBioStar) ToResponse() biostar.BioStarResponseDTO {
	return biostar.BioStarResponseDTO{
		UserId: mk.UserId,
		Name: mk.Name,
		UserGroupId: mk.UserGroupId,
		Disabled: mk.Disabled,
		StartDateTime: mk.StartDateTime,
		ExpiryDatetime: mk.ExpiryDatetime,
		Email: mk.Email,
		Department: mk.Department,
		Title: mk.Title,
		Photo: mk.Photo,
		Phone: mk.Phone,
		Permission: mk.Permission,
		AccessGroup: mk.AccessGroup,
		LoginID: mk.LoginID,
		Password: mk.Password,
		UserIp: mk.UserIp,
		Pin: mk.Pin,
	}
}