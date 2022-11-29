package views

import (
	"friendie/enums"
	"time"
)

type Pet struct {
	Id         int              `json:"id"`
	Name       string           `json:"name"`
	Birthday   time.Time        `json:"birthday"`
	Animal     enums.AnimalEnum `json:"animal"`
	Breed      enums.BreedEnum  `json:"breed"`
	FkAccount  int              `json:"fk_account"`
	Gender     enums.GenderEnum `json:"gender"`
	Sterilized bool             `json:"sterilized"`
}
