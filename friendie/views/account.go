package views

import (
	"friendie/enums"
	"time"
)

type Account struct {
	Id           int            `json:"id"`
	Name         string         `json:"name"`
	Surname      string         `json:"surname"`
	Birthday     time.Time      `json:"birthday"`
	EmailAddress string         `json:"email_address"`
	Password     string         `json:"password"`
	City         enums.CityEnum `json:"city"`
}
