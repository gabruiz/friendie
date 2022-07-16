package services

import (
	logic "backend/logics"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func addAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewUser")
	reqBody, _ := ioutil.ReadAll(r.Body)

	logic.AddAccount(reqBody)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/addAccount", addAccount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func StartRestService() {
	handleRequests()
}
