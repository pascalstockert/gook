package main

import (
	"encoding/json"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

type CronEntry struct {
	name    string
	spec    string
	closure func()
}

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
	cronEntries := []CronEntry{
		{
			name:    "test",
			spec:    "@every 5s",
			closure: func() { log.Println("test cron") },
		},
	}

	c := cron.New()

	for _, entry := range cronEntries {
		e := c.AddFunc(entry.spec, entry.closure)
		if e != nil {
			log.Fatalf("error adding cron entry: %v", e)
		}
	}

	c.Start()

	http.HandleFunc("/health", jsonMiddleware(healthHandler))

	log.Println(http.ListenAndServe(":4321", nil))
}
