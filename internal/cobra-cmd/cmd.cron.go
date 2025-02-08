package cobra_cmd

import (
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"go-webhook/pkg/env"
	"go-webhook/pkg/files"
	"go-webhook/pkg/types"
)

type cobraCommandFunc func(cmd *cobra.Command, args []string)

var CronAdd cobraCommandFunc = func(cmd *cobra.Command, args []string) {
	format := getFileFormat()
	parser, _ := files.GetParser(format)
	filePath := getFilePath(parser)
	entries := parser.ParseEntries(filePath)
	name, spec, protocol, location := getCronAddDialogResponses()

	entries = append(entries, types.CronEntry{
		Id:   uuid.New().String(),
		Name: name,
		Spec: spec,
		Action: types.CronAction{
			Type:     determineCronActionType(protocol),
			Resource: location,
		},
	})

	parser.WriteEntries(filePath, entries)
}

func determineCronActionType(protocol string) types.CronActionType {
	switch protocol {
	case "http":
		return types.CronActionTypeHttp
	}

	panic("Unsupported protocol: " + protocol)
}

func getFilePath(parser *types.FileParser) string {
	executableLocation, _ := os.Executable()
	pathArray := strings.Split(executableLocation, "/")
	workingDirectory := strings.Join(pathArray[:len(pathArray)-1], "/") + "/"
	fileName := "cron-entries"

	return workingDirectory + fileName + parser.FileSuffix
}

func getFileFormat() string {
	format, e := env.Get("CRON_FILE_FORMAT")
	if e != nil {
		panic(e)
	}

	return format
}
