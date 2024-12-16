package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonMiddleware(closure http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		closure(res, req)
	}
}

func healthHandler(res http.ResponseWriter, req *http.Request) {
	err := json.NewEncoder(res).Encode(`{status: ok}`)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	res.WriteHeader(http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/health", jsonMiddleware(healthHandler))

	log.Println(http.ListenAndServe(":4321", nil))
}
