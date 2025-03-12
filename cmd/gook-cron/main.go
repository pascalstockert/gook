package main

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
	"go-webhook/pkg/env"
	"go-webhook/pkg/files"
	"go-webhook/pkg/types"
)

var cronWorker *cron.Cron

func parseEntries() []types.CronEntry {
	format, e := env.Get("CRON_FILE_FORMAT")
	if e != nil {
		fmt.Println("could not get env-key CRON_FILE_FORMAT", e)
	}

	parser, e := files.GetParser(format)
	if e != nil {
		fmt.Println("could not get parser for format "+format, e)
	}

	filePath, e := parser.GetFilePath("cron-entries")
	if e != nil {
		fmt.Println("could not get file path for cron-entries", e)
	}

	return parser.ParseEntries(filePath)
}

var actionFunctionHttp = func(entry types.CronEntry) {
	req, err := http.NewRequest("GET", "http://"+entry.Action.Resource, nil)
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

func getActionFunctionForEntry(entry types.CronEntry) func() {
	functionMap := map[types.CronActionType]func(){
		types.CronActionTypeHttp: func() { actionFunctionHttp(entry) },
	}

	return functionMap[entry.Action.Type]
}

func addEntriesToCronWorker(entries []types.CronEntry) {
	// TODO store relation between cronjob and csv-entry to keep track of stopped/started jobs

	for _, entry := range entries {
		actionFunction := getActionFunctionForEntry(entry)

		err := cronWorker.AddFunc(entry.Spec, actionFunction)

		if err != nil {
			fmt.Println("could not add entry to cron worker", err)
		}
	}
}

func main() {
	env.Init()

	cronWorker = cron.New()
	cronWorker.Start()

	entries := parseEntries()
	addEntriesToCronWorker(entries)

	// keep alive
	<-make(chan struct{})
}
