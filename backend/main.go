package main

import (
	"backend/services"
	"log"
)

func main() {
	log.Output(0, "Hello Bonder")
	services.StartRestService()
}
