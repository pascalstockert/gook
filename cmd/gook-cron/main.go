package main

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
	"go-webhook/pkg/env"
	"go-webhook/pkg/files"
	"go-webhook/pkg/types"
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
			if entry.Action.Type == types.CronActionTypeHttp {
				req, err := http.NewRequest("GET", "https://"+entry.Action.Resource, nil)
				if err != nil {
					fmt.Println("could not create request", err)
					return
				}

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println("could not send request", err)
					return
				}

				fmt.Println(res.StatusCode)
			}

			// TODO determine action to execute
			fmt.Println(entry.Action)
		})
	}

	// keep alive
	<-make(chan struct{})
}
