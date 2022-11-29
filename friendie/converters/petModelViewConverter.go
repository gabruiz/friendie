package converters

import (
	"friendie/models"
	"friendie/views"
)

func PetCreateFromView(view views.Pet) models.Pet {
	var model models.Pet
	model.Id = view.Id
	model.Name = view.Name
	model.Birthday = view.Birthday
	model.Animal = view.Animal
	model.Breed = view.Breed
	model.FkAccount = view.FkAccount
	model.Gender = view.Gender
	model.Sterilized = view.Sterilized

	return model
}

func PetCreateFromModel(model models.Pet) views.Pet {
	var view views.Pet
	model.Name = view.Name
	model.Birthday = view.Birthday
	model.Animal = view.Animal
	model.Breed = view.Breed
	model.FkAccount = view.FkAccount
	model.Gender = view.Gender
	model.Sterilized = view.Sterilized

	return view
}
