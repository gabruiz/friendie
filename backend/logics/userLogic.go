package logic

import (
	"backend/db"
	"backend/views"
	"encoding/json"
	"log"
)

func AddAccount(reqBody []byte) {
	if reqBody != nil {
		var account views.AccountView
		err := json.Unmarshal(reqBody, &account)
		if err != nil {
			log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
		}

		db := db.DbConnection()
		ins, err := db.Prepare("INSERT INTO user (name, surname, email, password) VALUES (?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		ins.Exec(account.Name, account.Surname, account.Email, account.Password)
		defer db.Close()

	}
}
