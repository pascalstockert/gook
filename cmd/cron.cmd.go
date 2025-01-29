package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"go-webhook/shared/env"
	"go-webhook/shared/files"
	"go-webhook/shared/helper"
	"os"
)

type cobraCommandFunc func(cmd *cobra.Command, args []string)

var CronAdd cobraCommandFunc = func(cmd *cobra.Command, args []string) {
	var format, e = env.Get("CRON_FILE_FORMAT")
	if e != nil {
		panic(e)
	}

	var parser, _ = files.GetParser(format)
	entries := parser.ParseEntries("./cron-entries" + parser.FileSuffix)
	fmt.Println(entries)

	reader := bufio.NewReader(os.Stdin)

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

	name, spec, protocol, location := helper.Destructure4(responses)

	fmt.Println(name, spec, protocol, location)
}
