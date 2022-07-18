package builders

import (
	"bonder/models"
	"bonder/services/database"
)

var db = database.DatabaseConnection()

func SaveAccount(account models.AccountModel) models.AccountModel {
	ins, err := db.Prepare("INSERT INTO account (name, surname, email_address, password, city, creation_date) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(account.Name, account.Surname, account.EmailAddress, account.Password, account.City, account.CreationDate)
	defer db.Close()

	return account
}

func GetAccount(id int) models.AccountModel {
	ins, err := db.Prepare("SELECT * FROM account WHERE id = (?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(id)
	defer db.Close()

	return models.AccountModel{}
}
