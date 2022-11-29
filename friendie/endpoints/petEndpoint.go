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

func InitPetRestService(router *mux.Router) {
	router.HandleFunc("/pets", addPet).Methods(http.MethodPost)
	router.HandleFunc("/pets/{id:[0-9]+}", getPet).Methods(http.MethodGet)
	router.HandleFunc("/pets/{id:[0-9]+}", updatePet).Methods(http.MethodPut)
	router.HandleFunc("/pets/{id:[0-9]+}", deletePet).Methods(http.MethodDelete)
}

func addPet(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	if reqBody == nil {
		log.Fatalf("Input not valid")
	}

	var view views.Pet
	err := json.Unmarshal(reqBody, &view)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	view = logic.AddPet(view)
	backToTheFrontend(view, w)
}

func getPet(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	view := logic.GetPet(id)
	backToTheFrontend(view, w)
}

func updatePet(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	reqBody, _ := io.ReadAll(r.Body)
	if reqBody == nil {
		log.Fatalf("Input not valid")
	}

	var view views.Pet
	err := json.Unmarshal(reqBody, &view)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling. Error: %s", err.Error())
	}

	view = logic.UpdatePet(id, view)
	backToTheFrontend(view, w)
}

func deletePet(w http.ResponseWriter, r *http.Request) {
	stringId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(stringId)

	logic.DeletePet(id)
	var response = "Pet removed successfully"
	backToTheFrontend(response, w)
}
