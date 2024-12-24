package http

import (
	"encoding/json"
	"fmt"
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

func Execute(port string) {
	portPrepended := ":" + port
	http.HandleFunc("/health", jsonMiddleware(healthHandler))
	fmt.Println(http.ListenAndServe(portPrepended, nil))
}
