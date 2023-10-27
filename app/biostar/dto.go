package biostar

import "time"

type BioStarRequestDTO struct {
	Name        string `json:"name"`
	UserGroupId uint   `json:"usergroup_id"`
	LoginID     string `json:"loginid"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
}

type BioStarResponseDTO struct {
	UserId         string    `json:"userid"`
	Name           string    `json:"name"`
	UserGroupId    uint      `json:"usergroup_id"`
	Disabled       bool      `json:"disabled"`
	StartDateTime  time.Time  `json:"startdt"`
	ExpiryDatetime time.Time `json:"expirydt"`
	Email          string    `json:"email"`
	Department     string    `json:"department"`
	Title          string    `json:"title"`
	Photo          []byte    `json:"photo"`
	Phone          string    `json:"phone"`
	Permission     uint      `json:"permission"`
	AccessGroup    uint      `json:"accessgroup"`
	LoginID        string    `json:"loginid"`
	Password       string    `json:"password"`
	UserIp         string    `json:"userip"`
	Pin            uint      `json:"pin"`
}