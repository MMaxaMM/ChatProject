package main

import (
	"chat/internal/http-server/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", handlers.Chat).Methods("POST")
	router.HandleFunc("/history", handlers.History).Methods("POST")

	log.Println("Run LLM Server")
	err := http.ListenAndServe(":5050", router)
	log.Fatal(err)
}
