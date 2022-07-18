package endpoints

import (
	logic "bonder/logics"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

	view := logic.AddAccount(reqBody)
	backToTheFrontend(view, w)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: getAccount")
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	view := logic.GetAccount(id)
	backToTheFrontend(view, w)
}

func updateAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Endpoint Hit: updateAccount")
	reqBody, _ := ioutil.ReadAll(r.Body)

	view := logic.UpdateAccount(reqBody)
	backToTheFrontend(view, w)
}

func backToTheFrontend(view any, w http.ResponseWriter) {
	payload, err := json.Marshal(view)

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
