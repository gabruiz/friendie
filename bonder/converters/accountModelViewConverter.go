package converters

import (
	"bonder/models"
	"bonder/views"
)

func AccountCreateFromView(view views.AccountView) models.AccountModel {
	var model models.AccountModel
	model.Id = view.Id
	model.Name = view.Name
	model.Surname = view.Surname
	model.City = view.City
	model.Password = view.Password
	model.EmailAddress = view.EmailAddress
	model.CreationDate = view.CreationDate
	model.LastUpdateDate = view.LastUpdateDate

	return model
}

func AccountCreateFromModel(model models.AccountModel) views.AccountView {
	var view views.AccountView
	view.Id = model.Id
	view.Name = model.Name
	view.Surname = model.Surname
	view.City = model.City
	view.Password = model.Password
	view.EmailAddress = model.EmailAddress
	view.CreationDate = model.CreationDate
	view.LastUpdateDate = model.LastUpdateDate

	return view
}
