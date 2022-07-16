package services

import (
	logic "backend/logics"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func addAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewUser")
	reqBody, _ := ioutil.ReadAll(r.Body)

	logic.AddAccount(reqBody)
}

func accountHandleRequests() {
	log.Output(0, "Account Rest Service ready")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/accounts", addAccount).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}
