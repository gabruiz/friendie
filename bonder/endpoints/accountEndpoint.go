package endpoints

import (
	logic "bonder/logics"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func InitAccountRestService(router *mux.Router) {
	router.HandleFunc("/accounts", addAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id:[0-9]+}", getAccount).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{id:[0-9]+}", updateAccount).Methods(http.MethodPut)
}

func addAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: addAccount")
	reqBody, _ := ioutil.ReadAll(r.Body)

	logic.AddAccount(reqBody)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: getAccount")
	reqBody, _ := ioutil.ReadAll(r.Body)

	logic.GetAccount(reqBody)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: updateAccount")
	reqBody, _ := ioutil.ReadAll(r.Body)

	logic.UpdateAccount(reqBody)
}
