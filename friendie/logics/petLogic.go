package logic

import (
	"friendie/builders"
	"friendie/converters"
	"friendie/views"
	"time"
)

func AddPet(view views.Pet) views.Pet {
	model := converters.PetCreateFromView(view)
	model.CreateAt = time.Now()
	model = builders.SavePet(model)
	return converters.PetCreateFromModel(model)
}

func GetPet(id int) views.Pet {
	model := builders.GetPet(id)
	return converters.PetCreateFromModel(model)
}

func UpdatePet(id int, view views.Pet) views.Pet {
	model := builders.GetPet(id)

	if view.Name != "" {
		model.Name = view.Name
	}

	if !view.Birthday.IsZero() {
		model.Birthday = view.Birthday
	}

	if view.Animal != "" {
		model.Animal = view.Animal
	}

	if view.Breed != "" {
		model.Breed = view.Breed
	}

	model.Sterilized = view.Sterilized

	model.UpdateAt = time.Now()

	model = builders.UpdatePet(model)
	return converters.PetCreateFromModel(model)
}

func DeletePet(id int) {
	builders.DeletePet(id)
}
