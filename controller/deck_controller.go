package controller

import (
	"Toggl/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type DeckController interface {
	CreateDeck(w http.ResponseWriter, r *http.Request) error
	OpenDeck(w http.ResponseWriter, r *http.Request) error
	DrawCard(w http.ResponseWriter, r *http.Request) error
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create deck for id {}")

	shuffled, error := strconv.ParseBool(r.URL.Query().Get("shuffled"))
	if error != nil {
		shuffled = false
	}
	cards := r.URL.Query().Get("cards")
	slice := make([]string, 0)
	if len(cards) != 0 {
		slice = append(slice, strings.Split(cards, ",")...)
	}
	response, err := service.CreateDeck(shuffled, slice)
	if err != nil {
		http.Error(w, err.Message, err.ResponseCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func OpenDeck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Opening deck")

	params := mux.Vars(r)

	id := params["id"]
	response, err := service.OpenDeck(id)
	if err != nil {
		http.Error(w, err.Message, err.ResponseCode)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DrawCard(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Draw card")

	id := mux.Vars(r)["id"]
	count, error := strconv.Atoi(r.URL.Query().Get("count"))
	if error != nil {
		http.Error(w, error.Error(), 401)
		return
	}
	response, err := service.DrawCard(id, count)
	if err != nil {
		http.Error(w, err.Message, err.ResponseCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
