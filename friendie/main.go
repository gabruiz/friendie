package main

import (
	"friendie/services"
	"log"
)

func main() {
	err := log.Output(0, "Hello Friendie")
	if err != nil {
		return
	}
	services.StartRestService()
}
