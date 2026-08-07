package main

import (
	"gemgo/gemini"
	"gemgo/handlers"
	"gemgo/midleware"
	"log"
	"net/http"
)

func main() {

	geminiClient, err := gemini.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer geminiClient.Close()

	mux := http.NewServeMux()
	port := ":9090"

	mux.HandleFunc("/", handlers.HandlerFunc)
	mux.HandleFunc("/ask", handlers.MakeAskGemHandler(geminiClient))

	fileServer := http.FileServer(http.Dir("../static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	handler := midleware.LoggingMidleware(mux)
	handler = midleware.CORSMiddleware(handler)

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Error while trying to start serv:", err)
	}
}
