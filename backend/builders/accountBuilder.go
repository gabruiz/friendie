package builders

import (
	"backend/db"
	"backend/models"
)

func SaveAccount(account models.AccountModel) models.AccountModel {
	db := db.DbConnection()

	ins, err := db.Prepare("INSERT INTO account (name, surname, email_address, password, city, creation_date) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(account.Name, account.Surname, account.EmailAddress, account.Password, account.City, account.CreationDate)
	defer db.Close()

	return account
}
