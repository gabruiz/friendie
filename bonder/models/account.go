package models

import (
	"bonder/enums"
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
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
