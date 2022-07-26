package main

import (
	"Toggl/route"
	"log"
	"net/http"
)

func main() {
	router := route.AppRouter()

	log.Fatalln(http.ListenAndServe(":8080", router))
}
