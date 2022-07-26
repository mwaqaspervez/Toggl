package route

import (
	"Toggl/controller"
	"github.com/gorilla/mux"
)

//AppRouter registers all APIs.
func AppRouter() *mux.Router {
	router := mux.NewRouter()

	// Route handles & endpoints
	router.HandleFunc("/api/deck", controller.CreateDeck).Methods("POST")
	router.HandleFunc("/api/deck/{id}", controller.OpenDeck).Methods("GET")
	router.HandleFunc("/api/deck/{id}/draw", controller.DrawCard).Methods("PUT")

	return router
}
