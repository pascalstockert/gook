package main

import (
	"github.com/robfig/cron"
	"go-webhook/cmd"
	"log"
)

type CronEntry struct {
	name    string
	spec    string
	closure func()
}

func main() {
	cmd.Execute()

	c := cron.New()
	// c.Start()

	cronEntries := []CronEntry{
		{
			name:    "test",
			spec:    "@every 5s",
			closure: func() { log.Println("test cron") },
		},
	}

	for _, entry := range cronEntries {
		e := c.AddFunc(entry.spec, entry.closure)
		if e != nil {
			log.Fatalf("error adding cron entry: %v", e)
		}
	}
}
