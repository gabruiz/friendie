package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitMainRestService(router *mux.Router) {
	router.HandleFunc("/", homePage).Methods(http.MethodGet)
	router.HandleFunc("/ping", pong).Methods(http.MethodGet)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
	fmt.Println("Endpoint Hit: pong")
}

func mainHandleRequests(w http.ResponseWriter, r *http.Request, router *mux.Router) {
	homePage(w, r)
	router.HandleFunc("/ping", pong).Methods(http.MethodGet)
}