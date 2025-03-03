package cobra_cmd

import (
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
	filePath := parser.GetFilePath("cron-entries")
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

func getFileFormat() string {
	format, e := env.Get("CRON_FILE_FORMAT")
	if e != nil {
		panic(e)
	}

	return format
}
