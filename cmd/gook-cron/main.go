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

	for _, entry := range entries {
		fmt.Println(entry)
	}

	// keep alive
	<-make(chan struct{})
}
