package logic

import (
	"backend/builders"
	"backend/converters"
	"backend/views"
	"encoding/json"
	"log"
)

func AddAccount(reqBody []byte) views.AccountView {
	if reqBody == nil {
		log.Fatalf("Input not valid")
	}

	var account views.AccountView
	err := json.Unmarshal(reqBody, &account)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	var savedModel = builders.SaveAccount(converters.AccountCreateFromView(account))
	return converters.AccountCreateFromModel(savedModel)
}
