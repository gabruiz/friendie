package builders

import (
	"bonder/models"
	"bonder/services/database"
)

func SaveAccount(account models.AccountModel) models.AccountModel {
	db := database.DatabaseConnection()

	ins, err := db.Prepare("INSERT INTO account (name, surname, email_address, password, city, creation_date) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(account.Name, account.Surname, account.EmailAddress, account.Password, account.City, account.CreationDate)
	defer db.Close()

	return account
}
