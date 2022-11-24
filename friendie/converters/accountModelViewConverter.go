package converters

import (
	"friendie/models"
	"friendie/views"
)

func AccountCreateFromView(view views.Account) models.Account {
	var model models.Account
	model.Name = view.Name
	model.Surname = view.Surname
	model.City = view.City
	model.Password = view.Password
	model.EmailAddress = view.EmailAddress

	return model
}

func AccountCreateFromModel(model models.Account) views.Account {
	var view views.Account
	view.Id = model.Id
	view.Name = model.Name
	view.Surname = model.Surname
	view.City = model.City
	view.Password = model.Password
	view.EmailAddress = model.EmailAddress

	return view
}
