package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type MessageRole string

type Message struct {
	Role    MessageRole `json:"role"`
	Content string      `json:"content"`
}

func Generate(w http.ResponseWriter, r *http.Request) {
	prompt, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Fatal("Error in LLM Server")
	}

	log.Printf("Prompt: %s", string(prompt))

	time.Sleep(time.Second * 5)
	message := "Ответ ассистента..."

	res, err := json.Marshal(Message{Role: "assistent", Content: message})
	if err != nil {
		log.Fatal("Error in LLM Server")
	}

	w.Write(res)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", Generate).Methods("POST")

	log.Println("Run LLM Server")
	err := http.ListenAndServe(":5000", router)
	log.Fatal(err)
}
