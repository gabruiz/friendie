package endpoints

import (
	"encoding/json"
	logic "friendie/logics"
	"friendie/views"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InitAccountRestService(router *mux.Router) {
	router.HandleFunc("/accounts", addAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id:[0-9]+}", getAccount).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{id:[0-9]+}", updateAccount).Methods(http.MethodPut)
	router.HandleFunc("/accounts/deactivates/{id:[0-9]+}", deactivateAccount).Methods(http.MethodPut)
	router.HandleFunc("/accounts/activates/{id:[0-9]+}", activateAccount).Methods(http.MethodPut)

}

func addAccount(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	if reqBody == nil {
		log.Fatalf("Input not valid")
	}

	var view views.Account
	err := json.Unmarshal(reqBody, &view)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	view = logic.AddAccount(view)
	backToTheFrontend(view, w)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	view := logic.GetAccount(id)
	backToTheFrontend(view, w)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	reqBody, _ := io.ReadAll(r.Body)
	if reqBody == nil {
		log.Fatalf("Input not valid")
	}

	var view views.Account
	err := json.Unmarshal(reqBody, &view)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	view = logic.UpdateAccount(id, view)
	backToTheFrontend(view, w)
}

func deactivateAccount(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	logic.DeactivateAccount(id)
	var response = "Account disattivato con successo"
	backToTheFrontend(response, w)
}

func activateAccount(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	logic.DeactivateAccount(id)
	var response = "Account removed successfully"
	backToTheFrontend(response, w)
}
