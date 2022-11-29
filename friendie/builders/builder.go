package builders

import (
	"friendie/services/database"

	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	return database.Connection()
}

func checkErrors(result *gorm.DB) {
	if result.Error != nil {
		panic(result.Error.Error())
	}
}
