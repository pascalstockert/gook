package main

import (
	"github.com/robfig/cron"
	"go-webhook/shared/types"
	"log"
)

func main() {
	c := cron.New()
	c.Start()

	// TODO read cronEntries from file
	cronEntries := []types.CronEntry{
		{
			Id:      "123",
			Name:    "test",
			Spec:    "@every 5s",
			Closure: func() { log.Println("test cron") },
		},
	}

	for _, entry := range cronEntries {
		e := c.AddFunc(entry.Spec, entry.Closure)
		if e != nil {
			log.Fatalf("error adding cron entry: %v", e)
		}
	}

	<-make(chan struct{})
}
