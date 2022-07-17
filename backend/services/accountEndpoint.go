package services

import (
	logic "backend/logics"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func addAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: createNewUser")
	reqBody, _ := ioutil.ReadAll(r.Body)

	logic.AddAccount(reqBody)
}

func accountHandleRequests(w http.ResponseWriter, r *http.Request, router *mux.Router) {
	fmt.Println(w, "Welcome to the account rest service")
	addAccount(w, r)
}
