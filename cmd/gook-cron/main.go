package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/robfig/cron"
	"go-webhook/pkg/env"
	"go-webhook/pkg/files"
	"go-webhook/pkg/types"
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

	filePath := getFilePath(parser)

	entries := parser.ParseEntries(filePath)

	for _, entry := range entries {
		fmt.Println(entry)
	}

	// keep alive
	<-make(chan struct{})
}

func getFilePath(parser *types.FileParser) string {
	executableLocation, _ := os.Executable()
	pathArray := strings.Split(executableLocation, "/")
	workingDirectory := strings.Join(pathArray[:len(pathArray)-1], "/") + "/"
	fileName := "cron-entries"

	return workingDirectory + fileName + parser.FileSuffix
}
