package builders

import (
	"friendie/models"
	"friendie/services/database"
)

var db = database.Connection()

func SaveAccount(model models.Account) models.Account {
	ins, err := db.Prepare("INSERT INTO accounts (name, surname, email_address, password, city, created_at) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(model.Name, model.Surname, model.EmailAddress, model.Password, model.City, model.CreatedAt)

	return model
}

func GetAccount(id int) models.Account {
	var model models.Account
	results, err := db.Query("SELECT id, name, surname, email_address, city FROM accounts WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		err = results.Scan(&model.Id, &model.Name, &model.Surname, &model.EmailAddress, &model.City)
		if err != nil {
			panic(err.Error())
		}
	}

	return model
}

func UpdateAccount(model models.Account) models.Account {
	_, err := db.Query("UPDATE accounts SET name ?, surname ?, city ?, updated_at ? where id = ?", model.Name, model.Surname, model.City, model.UpdatedAt, model.Id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return model
}
