package views

import (
	"friendie/enums"
)

type Account struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	Surname      string         `json:"surname"`
	EmailAddress string         `json:"email_address"`
	Password     string         `json:"password"`
	City         enums.CityEnum `json:"city"`
}
