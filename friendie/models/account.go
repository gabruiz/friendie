package models

import (
	"friendie/enums"
	"time"
)

type Account struct {
	Id           int
	Name         string
	Surname      string
	EmailAddress string
	Password     string
	City         enums.CityEnum
	Active       bool
	CreateAt     time.Time
	UpdateAt     time.Time
}
