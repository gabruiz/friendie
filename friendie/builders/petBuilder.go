package builders

import (
	"friendie/models"
)

func SavePet(model models.Pet) models.Pet {
	result := getDb().FirstOrCreate(&model)
	checkErrors(result)

	return model
}

func GetPet(id int) models.Pet {
	var model models.Pet
	result := getDb().First(&model, id)
	checkErrors(result)

	return model
}

func UpdatePet(model models.Pet) models.Pet {
	result := getDb().First(&model)
	checkErrors(result)

	return model
}

func DeletePet(id int) {
	result := getDb().Delete(nil, id)
	checkErrors(result)
}
