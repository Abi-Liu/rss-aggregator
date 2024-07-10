package main

import "net/http"

func getHealthStatus(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}

	respondWithJson(w, http.StatusOK, response{
		Status: "ok",
	})
}

func simulateError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
