package models

import (
	"friendie/enums"
	"time"
)

type Pet struct {
	Id         int
	Name       string
	Birthday   time.Time
	Animal     enums.AnimalEnum
	Breed      enums.BreedEnum
	FkAccount  int
	Gender     enums.GenderEnum
	Sterilized bool
	CreateAt   time.Time
	UpdateAt   time.Time
}
