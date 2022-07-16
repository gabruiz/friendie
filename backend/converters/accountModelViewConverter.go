package converters

import (
	"backend/models"
	"backend/views"
)

func CreateFromView(views.AccountView) models.AccountModel {
	return models.AccountModel{}
}
