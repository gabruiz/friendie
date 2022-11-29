package builders

import (
	"friendie/models"
)

func SaveAccount(model models.Account) models.Account {
	result := getDb().FirstOrCreate(&model)
	checkErrors(result)

	return model
}

func GetAccount(id int) models.Account {
	var model models.Account
	result := getDb().First(&model, id)
	checkErrors(result)

	return model
}

func UpdateAccount(model models.Account) models.Account {
	result := getDb().First(&model)
	checkErrors(result)

	return model
}

func DeactivateAccount(id int) {
	result := getDb().Model(nil).Where("id = ?", id).Update("active", false)
	checkErrors(result)
}

func ActivateAccount(id int) {
	result := getDb().Model(nil).Where("id = ?", id).Update("active", true)
	checkErrors(result)
}
