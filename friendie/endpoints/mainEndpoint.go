package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitMainRestService(router *mux.Router) {
	router.HandleFunc("/", homePage).Methods(http.MethodGet)
	router.HandleFunc("/ping", pong).Methods(http.MethodGet)
}

func homePage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func pong(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Pong!")
	fmt.Println("Endpoint Hit: pong")
}

func backToTheFrontend(view any, w http.ResponseWriter) {
	payload, err := json.Marshal(view)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
