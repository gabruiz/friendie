package services

import "log"

func StartRestService() {
	log.Output(0, "Loading Rest Service")
	go accountHandleRequests()
	go mainHandleRequests()
}
