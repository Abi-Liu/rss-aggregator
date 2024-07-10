package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json: %s\n", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(status)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, status int, msg string) {
	if status > 499 {
		log.Printf("Internal error: %s\n", msg)
	}

	type response struct {
		Error string `json:"error"`
	}

	respondWithJson(w, status, response{
		Error: msg,
	})
}
