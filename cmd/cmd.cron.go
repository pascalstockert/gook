package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"go-webhook/shared/env"
	"go-webhook/shared/files"
	"go-webhook/shared/helper"
	"go-webhook/shared/types"
)

type cobraCommandFunc func(cmd *cobra.Command, args []string)

var CronAdd cobraCommandFunc = func(cmd *cobra.Command, args []string) {
	format, e := env.Get("CRON_FILE_FORMAT")
	if e != nil {
		panic(e)
	}

	parser, _ := files.GetParser(format)

	// TODO make filePath absolute with os.Executable()
	filePath := "./cron-entries" + parser.FileSuffix
	entries := parser.ParseEntries("./cron-entries" + parser.FileSuffix)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(entries)

	responses := Dialog(*reader, []Phrase{
		{
			Claim: "Name the cronjob:",
		},
		{
			Claim: "Specify the cron-spec",
		},
		{
			Claim: "Choose a protocol:",
			Options: []string{
				"http",
			},
		},
		{
			Claim: "Specify the resource location:",
		},
	})

	determineCronActionType := func(protocol string) types.CronActionType {
		switch protocol {
		case "http":
			return types.CronActionTypeHttp
		}

		panic("Unsupported protocol: " + protocol)
	}

	name, spec, protocol, location := helper.Destructure4(responses)

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

	fmt.Println(entries)
}
