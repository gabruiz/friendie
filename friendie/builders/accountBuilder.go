package builders

import (
	"friendie/models"
	"friendie/services/database"
)

var db = database.Connection()

func SaveAccount(model models.Account) models.Account {
	ins, err := db.Prepare("INSERT INTO accounts (name, surname, email_address, password, city, create_at) VALUES (?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ins.Exec(model.Name, model.Surname, model.EmailAddress, model.Password, model.City, model.CreateAt)

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
	_, err := db.Query("UPDATE accounts SET name ?, surname ?, city ?, update_at ? where id = ?", model.Name, model.Surname, model.City, model.UpdateAt, model.Id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return model
}

func DeactivateAccount(id int) {
	_, err := db.Query("UPDATE accounts SET active FALSE where id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

func ActivateAccount(id int) {
	_, err := db.Query("UPDATE accounts SET active TRUE where id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
