package main

import (
	"fmt"

	"github.com/robfig/cron"
	"go-webhook/pkg/env"
	"go-webhook/pkg/files"
)

func main() {
	env.Init()

	c := cron.New()
	c.Start()

	// TODO refactor into generic function & call once for initiation
	format, e := env.Get("CRON_FILE_FORMAT")
	if e != nil {
		fmt.Println("could not get env", e)
	}

	parser, e := files.GetParser(format)
	if e != nil {
		fmt.Println("could not get parser", e)
	}

	filePath := parser.GetFilePath("cron-entries")

	entries := parser.ParseEntries(filePath)

	// TODO store relation between cronjob and csv-entry to keep track of stopped/started jobs
	for _, entry := range entries {
		_ = c.AddFunc(entry.Spec, func() {
			// TODO determine action to execute
			fmt.Println(entry.Action)
		})
	}

	// keep alive
	<-make(chan struct{})
}
