package logic

import (
	"friendie/builders"
	"friendie/converters"
	"friendie/views"
	"time"
)

func AddAccount(view views.Account) views.Account {
	model := converters.AccountCreateFromView(view)
	model.CreatedAt = time.Now()
	model = builders.SaveAccount(model)
	return converters.AccountCreateFromModel(model)
}

func GetAccount(id int) views.Account {
	model := builders.GetAccount(id)
	return converters.AccountCreateFromModel(model)
}

func UpdateAccount(id int, view views.Account) views.Account {
	model := builders.GetAccount(id)

	if view.Name != "" {
		model.Name = view.Name
	}

	if view.Surname != "" {
		model.Surname = view.Surname
	}

	if view.EmailAddress != "" {
		model.EmailAddress = view.EmailAddress
	}

	if view.Password != "" {
		model.Password = view.Password
	}

	model.UpdatedAt = time.Now()

	model = builders.UpdateAccount(model)
	return converters.AccountCreateFromModel(model)
}
