package main

import (
	"log"

	"github.com/robfig/cron"
	"go-webhook/shared/types"
)

func main() {
	c := cron.New()
	c.Start()

	// TODO read cronEntries from file
	cronEntries := []types.CronEntry{
		{
			Id:   "1",
			Name: "pasu-home",
			Spec: "@every 5s",
			Action: types.CronAction{
				Type:     "http",
				Resource: "https://pasu.me",
			},
		},
	}

	for _, entry := range cronEntries {
		// TODO determine function to execute from CronEntry.Action
		e := c.AddFunc(entry.Spec, func() {})
		if e != nil {
			log.Fatalf("error adding cron entry: %v", e)
		}
	}

	<-make(chan struct{})
}
