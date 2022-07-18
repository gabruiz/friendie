package main

import (
	"bonder/services"
	"log"
)

func main() {
	log.Output(0, "Hello Bonder")
	services.StartRestService()
}
