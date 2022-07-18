package services

import (
	"bonder/endpoints"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRestService() {
	log.Println("Loading Rest Service")

	router := mux.NewRouter()
	endpoints.InitAccountRestService(router)
	endpoints.InitMainRestService(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8083", router))
}
