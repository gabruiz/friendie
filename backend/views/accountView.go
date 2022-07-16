package views

import (
	"backend/enums"
	"time"
)

type AccountView struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	Surname        string         `json:"surname"`
	EmailAddress   string         `json:"email_address"`
	Password       string         `json:"password"`
	City           enums.CityEnum `json:"city"`
	CreationDate   time.Time      `json:"creation_date"`
	LastUpdateDate time.Time      `json:"last_update_date"`
	// active         bool           `json:"active"`
}
