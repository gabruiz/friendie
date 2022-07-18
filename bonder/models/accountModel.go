package models

import "time"
import "bonder/enums"

type AccountModel struct {
	Id             int
	Name           string
	Surname        string
	EmailAddress   string
	Password       string
	City           enums.CityEnum
	CreationDate   time.Time
	LastUpdateDate time.Time
	active         bool
}
