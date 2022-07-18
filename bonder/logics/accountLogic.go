package logic

import (
	"bonder/builders"
	"bonder/converters"
	"bonder/views"
)

func AddAccount(account views.AccountView) views.AccountView {
	var savedModel = builders.SaveAccount(converters.AccountCreateFromView(account))
	return converters.AccountCreateFromModel(savedModel)
}

func GetAccount(id int) views.AccountView {
	model := builders.GetAccount(id)
	return converters.AccountCreateFromModel(model)
}

func UpdateAccount(id int, account views.AccountView) views.AccountView {
	return views.AccountView{}
}
