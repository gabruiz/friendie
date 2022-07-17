package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartRestService() {
	log.Println("Loading Rest Service")

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { mainHandleRequests(w, r, router) }).Methods(http.MethodGet)
	router.HandleFunc("/accounts", func(w http.ResponseWriter, r *http.Request) { accountHandleRequests(w, r, router) }).Methods(http.MethodPost)
	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8083", router))
}
