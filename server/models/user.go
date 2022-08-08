package models

import "myzavod/pkg/tools"

// User ...
type User struct {
	tools.Model

	Email             string     `gorm:"type:varchar(100);unique" json:"email"`
	Name              string     `gorm:"type:varchar(255)" json:"name"`
	Phone             string     `gorm:"type:varchar(30)" json:"phone"`
	EmailVerified     bool       `json:"email_verified"`
	PayService        string     `json:"pay_service,omitempty"`
	PreferredLocation *Location  `json:"preferred_location,omitempty"`
	Locations         []Location `json:"locations,omitempty"`

	IsAdmin           bool   `json:"is_admin,omitempty"`
	PasswordHash      string `gorm:"type:varchar(64)" json:"-"`
	AlterPasswordHash string `gorm:"type:varchar(64)" json:"-"`
}

// Location ...
type Location struct {
	tools.Model

	UserID       uint   `json:"user_id"`
	Country      string `json:"country"`
	DistrictCity string `json:"district_city"`
	StreetHome   string `json:"street_home"`
}
