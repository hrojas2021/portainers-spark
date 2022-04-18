package app

import (
	"encoding/json"
	"log"
	"net/http"
	"spark/challenge/portainer"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Println("Listen port" + addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/validate/{setup}", a.validate).Methods("GET")
}

func (a *App) validate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	isValidSetup := portainer.ValidateSetup(params["setup"])
	if !isValidSetup {
		respondWithError(w, http.StatusBadRequest, "Invalid Setup params")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Valid Setup!"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
